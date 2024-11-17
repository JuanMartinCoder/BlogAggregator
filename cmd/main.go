package main

import (
	"log"
	"os"

	"github.com/JuanMartinCoder/BlogAggregator/cmd/commands"
	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	programState := config.State{Cfg: &cfg}

	cmds := commands.CreateCommands()

	cmds.Register("login", commands.LoginHandler)

	if len(os.Args) < 2 {
		log.Fatal("No command provided")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(&programState, commands.CliCommand{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
