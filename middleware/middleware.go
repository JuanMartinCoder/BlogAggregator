package middleware

import (
	"context"
	"fmt"

	"github.com/JuanMartinCoder/BlogAggregator/cmd/commands"
	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *config.State, cmd commands.CliCommand, user database.User) error) func(*config.State, commands.CliCommand) error {
	return func(s *config.State, cmd commands.CliCommand) error {
		user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
