package commands

import (
	"context"
	"fmt"

	"github.com/JuanMartinCoder/BlogAggregator/internal/config"
	"github.com/JuanMartinCoder/BlogAggregator/internal/rssfeed"
)

func AggHandler(s *config.State, cmd CliCommand) error {
	rssfeed, err := rssfeed.FetchRSSFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch rss feed: %w", err)
	}
	fmt.Println(rssfeed)
	return nil
}
