package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Throne-of-Doom/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username required")
	}
	username := cmd.args[0]
	param := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(ctx, param)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return fmt.Errorf("user %q already exists", username)
		}
		return fmt.Errorf("create user: %w", err)
	}
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("set user: %w", err)
	}
	fmt.Printf("created user: %+v\n", user)
	fmt.Println("user has been set")
	return nil
}
