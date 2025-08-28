package main

import (
	"fmt"
	"os"

	"github.com/CodyMcCarty/aggreGATOR/internal/config"
)

type state struct {
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
	lCfg, err := config.Read()
	if err != nil {
		return err
	}
	prevName := lCfg.CurrentUserName

	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("UserNameChanged %s -> %s\n", prevName, cmd.args[0])

	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	s := &state{
		config: &cfg,
	}

	cmds := &commands{
		commandFuncMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

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
