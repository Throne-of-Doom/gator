package main

import (
	"fmt"

	"github.com/Throne-of-Doom/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't auto-follow new feed: %w", err)
	}
	printFeed(feed)
	return nil
}

// go
func printFeed(feed database.Feed) {
	fmt.Printf("* ID:      %s\n", feed.ID)
	fmt.Printf("* Created: %v\n", feed.CreatedAt)
	fmt.Printf("* Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:    %s\n", feed.Name)
	fmt.Printf("* URL:     %s\n", feed.Url)
	fmt.Printf("* UserID:  %s\n", feed.UserID)
}
