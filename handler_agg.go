package main

import (
	"database/sql"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: Agg <time>")
	}
	time_between_req := cmd.args[0]
	newTime, err := time.ParseDuration(time_between_req)
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}
	fmt.Printf("Collecting feeds every %v", newTime)
	ticker := time.NewTicker(newTime)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(ctx)
	if err == sql.ErrNoRows {
		fmt.Println("No feeds to fetch")
		return nil
	}
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(ctx, next.ID)
	if err != nil {
		return fmt.Errorf("error marking feed %d fetched: %w", next.ID, err)
	}
	feed, err := fetchFeed(ctx, next.Url)
	if err != nil {
		return fmt.Errorf("error getting feed at %s: %w", next.Url, err)
	}
	for _, post := range feed.Channel.Item {
		fmt.Println("Post:", post.Title)
	}

	return nil
}
