package commands

import (
	"errors"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
)

type CliCommand struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*config.State, CliCommand) error
}

func (c *Commands) Register(name string, f func(*config.State, CliCommand) error) {
	c.RegisteredCommands[name] = f
}

func (c *Commands) Run(s *config.State, cmd CliCommand) error {
	f, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}
	return f(s, cmd)
}

func CreateCommands() *Commands {
	return &Commands{
		RegisteredCommands: make(map[string]func(*config.State, CliCommand) error),
	}
}
