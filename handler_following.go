package main

import (
	"fmt"

	"github.com/Throne-of-Doom/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: following")
	}
	rows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}
	for _, row := range rows {
		fmt.Println(row.FeedName)
	}
	return nil
}
