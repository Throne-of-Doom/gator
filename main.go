package main

import (
	"fmt"
	"github.com/Throne-of-Doom/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("Chris")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	fmt.Println(cfg)
}
