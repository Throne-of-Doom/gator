package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username required")
	}
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user %q doesn't exist", name)
	}
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Println("user has been set")
	return s.cfg.SetUser(cmd.args[0])
}
