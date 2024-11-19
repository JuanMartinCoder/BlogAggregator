package commands

import (
	"context"
	"fmt"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
)

func UsersHandler(s *config.State, cmd CliCommand) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
