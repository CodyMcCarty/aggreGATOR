package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/CodyMcCarty/aggreGATOR/internal/database"
	_ "github.com/lib/pq"

	"github.com/CodyMcCarty/aggreGATOR/internal/config"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Open a connection to the database, and store it in the state struct:
	// In main(), load in your database URL to the config struct and sql.Open() a connection to your database:
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	// Use your generated database package to create a new *database.Queries, and store it in your state struct:
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	// Create a register handler and register it with the commands. Usage: `go run . register lane`
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
