package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/austinthieu/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	// default value for limit
	limit := 2

	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("couldn't convert limit string to int: %w", err)
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostFromUser(context.Background(), database.GetPostFromUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts from user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
