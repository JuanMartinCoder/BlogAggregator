package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JuanMartinCoder/BlogAggregator/cmd/commands"
	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Database
	db, err := sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &config.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}

	cmds := commands.CreateCommands()

	cmds.Register("login", commands.LoginHandler)
	cmds.Register("register", commands.RegisterHandler)
	cmds.Register("reset", commands.ResetHandler)
	cmds.Register("users", commands.UsersHandler)
	cmds.Register("agg", commands.AggHandler)
	cmds.Register("addfeed", commands.FeedHandler)
	cmds.Register("feeds", commands.FeedsHandler)

	if len(os.Args) < 2 {
		log.Fatal("No command provided")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, commands.CliCommand{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
