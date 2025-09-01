package main

import (
	"context"
	"fmt"
	"time"

	"github.com/austinthieu/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed Follow successfully created!")

	printFeedFollows(feedFollow, user, feed)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s unfolowed successfully!\n", feed.Name)

	return nil
}

func printFeedFollows(ff database.CreateFeedFollowsRow, user database.User, feed database.Feed) {
	fmt.Printf("* ID:               %s\n", ff.ID)
	fmt.Printf("* CreatedAt:        %s\n", ff.CreatedAt)
	fmt.Printf("* UpdatedAt:        %s\n", ff.UpdatedAt)
	fmt.Printf("* Feed:             %s\n", feed.Name)
	fmt.Printf("* User:             %s\n", user.Name)
	fmt.Println("==================================================")
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("\n%s is following these feeds\n", user.Name)
	fmt.Println("==================================================")

	for _, f := range feedFollows {
		fmt.Printf("Name: %s\n", f.FeedName)
		fmt.Println("==================================================")
	}

	return nil
}
