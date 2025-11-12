package main

import (
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	results, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	fmt.Println(results)
	return nil
}
