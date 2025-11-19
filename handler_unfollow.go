package main

import (
	"fmt"

	"github.com/Throne-of-Doom/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	url := cmd.args[0]
	_, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("feed not found for url %q: %w", url, err)
	}
	_, err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}
	fmt.Println("Feed unfollowed successfully!")
	return nil
}
