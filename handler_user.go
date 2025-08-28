package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/CodyMcCarty/aggreGATOR/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Update the login command handler to error (and exit with code 1) if the given username doesn't exist in the database. "You can't login to an account that doesn't exist!"
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("You can't login to an account that doesn't exist!", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	// Ensure that a name was passed in the args.
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	name := strings.TrimSpace(cmd.Args[0])
	if name == "" {
		return fmt.Errorf("username is required")
	}

	// Exit with code 1 if a user with that name already exists.
	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		return fmt.Errorf("username '%q' is already taken", name)
	}

	// Create a new user in the database. It should have access to the CreateUser query through the state -> db struct.
	// Use the uuid.New() function to generate a new UUID for the user.
	// created_at and updated_at should be the current time.
	// Use the provided name.
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	// Pass context.Background() to the query to create an empty Context argument.
	usr, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	// Set the current user in the config to the given name.
	if err := s.cfg.SetUser(usr.Name); err != nil {
		return err
	}
	s.cfg.CurrentUserName = usr.Name

	// Print a message that the user was created, and log the user's data to the console for your own debugging.
	fmt.Println(s.cfg.CurrentUserName, "was created")
	fmt.Printf("debug: %+v\n", usr)

	return nil
}
