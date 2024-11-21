package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func FollowHandler(s *config.State, cmd CliCommand, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	feedId, err := s.DB.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedId.ID,
	}

	records, err := s.DB.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println(records)

	return nil
}

func FollowingHandler(s *config.State, cmd CliCommand, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	for _, follow := range follows {
		fmt.Printf("Feed: %s, URL: %s Created By: %s\n", follow.FeedName, follow.FeedName, follow.UserName)
	}

	return nil
}

func UnfollowHandler(s *config.State, cmd CliCommand, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	feedId, err := s.DB.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.DB.UnFollowFeed(context.Background(), database.UnFollowFeedParams{
		UserID: user.ID,
		FeedID: feedId.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}

	return nil
}
