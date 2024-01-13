package runic

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

type NewznabReader struct{}

func (r *NewznabReader) ReadURL(url string) ([]*Item, error) {
	feed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}
	return r.process(feed)
}

func (r *NewznabReader) ReadFile(path string) ([]*Item, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	feed, err := gofeed.NewParser().Parse(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}
	return r.process(feed)
}

func (r *NewznabReader) process(feed *gofeed.Feed) ([]*Item, error) {
	out := []*Item{}
	for _, i := range feed.Items {
		item := &Item{
			Title:       i.Title,
			Description: i.Description,
		}
		if i.PublishedParsed != nil {
			item.Published = *i.PublishedParsed
		}
		out = append(out, item)
	}
	return out, nil
}
