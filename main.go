package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Throne-of-Doom/gator/internal/config"
	"github.com/Throne-of-Doom/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Get database URL from config
	dbURL := cfg.DbURL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	c := &commands{}
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	name := args[1]
	argsSlice := args[2:]
	cmd := command{name: name, args: argsSlice}

	if err := c.run(s, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
