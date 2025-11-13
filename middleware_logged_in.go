package main

import (
	"fmt"

	"github.com/Throne-of-Doom/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		return handler(s, cmd, user)
	}
}
