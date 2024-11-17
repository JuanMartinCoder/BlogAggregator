package commands

import (
	"fmt"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
)

func LoginHandler(s *config.State, cmd CliCommand) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User switched successfully")

	return nil
}
