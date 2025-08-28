package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CodyMcCarty/aggreGATOR/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/CodyMcCarty/aggreGATOR/internal/config"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandFuncMap map[string]func(*state, command) error
}

// Runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandFuncMap[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	return f(s, cmd)
}

// Registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.commandFuncMap[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	usr, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	if usr.Name == "" {
		return fmt.Errorf("You can't login to an account that doesn't exist!")
	}

	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("UserNameChanged-> %s\n", cmd.args[0])

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	name := strings.TrimSpace(cmd.args[0])
	if name == "" {
		return fmt.Errorf("username is required")
	}

	// Exit with code 1 if a user with that name already exists.
	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		return fmt.Errorf("username '%q' is already taken", name)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("checking existing user: %w", err)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	usr, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	s.config.CurrentUserName = usr.Name

	fmt.Println(s.config.CurrentUserName, "was created")
	fmt.Printf("debug: %+v\n", usr)

	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := &commands{
		commandFuncMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	err = cmds.run(s, command{cmdName, cmdArgs})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
