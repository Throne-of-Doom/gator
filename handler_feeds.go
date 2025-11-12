package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	rows, err := s.db.ListFeeds(ctx)
	if err != nil {
		return err
	}
	for _, row := range rows {
		fmt.Println(row.FeedName)
		fmt.Println(row.FeedUrl)
		fmt.Println(row.UserName)
	}
	return nil
}
