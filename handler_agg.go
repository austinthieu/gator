package main

import (
	"context"
	"fmt"

	"github.com/austinthieu/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func addFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("need name and url in arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	userId, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: userId.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	fmt.Println(feed)

	return nil
}
