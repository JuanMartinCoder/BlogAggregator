package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JuanMartinCoder/BlogAggregator/cmd/commands"
	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
	"github.com/JuanMartinCoder/BlogAggregator/middleware"

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
	cmds.Register("addfeed", middleware.MiddlewareLoggedIn(commands.FeedHandler))
	cmds.Register("feeds", commands.FeedsHandler)
	cmds.Register("follow", middleware.MiddlewareLoggedIn(commands.FollowHandler))
	cmds.Register("following", middleware.MiddlewareLoggedIn(commands.FollowingHandler))
	cmds.Register("unfollow", middleware.MiddlewareLoggedIn(commands.UnfollowHandler))
	cmds.Register("browse", middleware.MiddlewareLoggedIn(commands.BrowseHandler))

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
