package main

import (
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	fmt.Println("users has been reset")
	return nil
}
