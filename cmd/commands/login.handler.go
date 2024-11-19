package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func LoginHandler(s *config.State, cmd CliCommand) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	_, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	err = s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println("User switched successfully")

	return nil
}

func RegisterHandler(s *config.State, cmd CliCommand) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	arg := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}

	user, err := s.DB.CreateUser(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Printf("User: %v+", user)
	fmt.Println("User created successfully")

	return nil
}
