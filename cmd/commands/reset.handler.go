package commands

import (
	"context"
	"fmt"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
)

// Only for testing purposes, don't use in production
func ResetHandler(s *config.State, cmd CliCommand) error {
	err := s.DB.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}
	fmt.Println("Users reset successfully")
	return nil
}
