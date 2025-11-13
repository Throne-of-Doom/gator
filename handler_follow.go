package main

import (
	"fmt"

	"github.com/Throne-of-Doom/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("feed not found for url %q: %w", url, err)
	}

	rows, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}
	if len(rows) == 0 {
		return fmt.Errorf("no follow row returned")
	}
	printFollow(rows[0])
	return nil
}

func printFollow(row database.CreateFeedFollowRow) {
	fmt.Printf("%s follows %s\n", row.UserName, row.FeedName)
}
