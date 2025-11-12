package main

import (
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("get users: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		line := "* " + user
		if user == currentUser {
			line += " (current) "
		}
		fmt.Println(line)

	}
	return nil
}
