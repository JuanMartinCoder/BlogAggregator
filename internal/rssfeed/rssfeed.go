package rssfeed

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

func FetchRSSFeed(ctx context.Context, url string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &RSSFeed{}, err
	}

	body, err := io.ReadAll(resp.Body)
	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return &RSSFeed{}, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, Item := range feed.Channel.Item {
		Item.Title = html.UnescapeString(Item.Title)
		Item.Description = html.UnescapeString(Item.Description)
		feed.Channel.Item[i] = Item
	}

	return &feed, nil
}
