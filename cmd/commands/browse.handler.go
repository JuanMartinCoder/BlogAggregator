package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
)

func BrowseHandler(s *config.State, cmd CliCommand, user database.User) error {
	var limit int
	if len(cmd.Args) != 1 {
		limit = 2
	} else {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("couldn't parse limit: %w", err)
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Post: %s, URL: %s, Description: %s, Published At: %s\n", post.Title, post.Url, post.Description, post.PublishedAt)
	}

	return nil
}
