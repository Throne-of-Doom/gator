package main

import (
	"fmt"
	"github.com/Throne-of-Doom/gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := &state{cfg: &cfg}
	c := &commands{}
	c.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	name := args[1]
	argsSlice := args[2:]
	cmd := command{name: name, args: argsSlice}
	if err := c.run(s, cmd); err != nil {
		fmt.Println("username is required")
		os.Exit(1)
	}
}
