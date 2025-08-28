package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("Username does not exist: %s", username)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("Username has been set successfully")

	return nil
}
