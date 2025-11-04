package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username required")
	}
	fmt.Println("user has been set")
	return s.cfg.SetUser(cmd.args[0])
}
