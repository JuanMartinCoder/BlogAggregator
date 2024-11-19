package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func FeedHandler(s *config.State, cmd CliCommand) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	actualUser, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	arg := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    actualUser.ID,
	}

	_, err = s.DB.CreateFeed(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	return nil
}

func FeedsHandler(s *config.State, cmd CliCommand) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s, URL: %s Created By: %s\n", feed.Name, feed.Url, feed.Name_2)
	}

	return nil
}
