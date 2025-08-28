package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/austinthieu/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), name)

	if user.Name != "" {
		fmt.Printf("error: User with the name '%s' already exists.\n", name)
		os.Exit(1)
	}
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	fmt.Println(user.Name)
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User successfully created!")

	return nil
}
