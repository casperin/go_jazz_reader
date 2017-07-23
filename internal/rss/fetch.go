package rss

import (
	"fmt"
	"time"

	"github.com/casperin/jazz_reader/internal/feed"
	"github.com/casperin/jazz_reader/internal/post"
	"github.com/casperin/jazz_reader/internal/util/hash"
	"github.com/casperin/jazz_reader/internal/util/str"
	"github.com/mmcdole/gofeed"
)

func FetchUrl(url string) (*feed.Feed, []*post.Post, error) {
	fetched, err := gofeed.NewParser().ParseURL(url)
	if fetched == nil {
		return nil, nil, fmt.Errorf("Could not parse %s:\n%v", url, err)
	}
	f := parseFeed(url, fetched, err)
	ps := parsePosts(fetched.Items, f.Title)
	return f, ps, err
}

func parseFeed(url string, f *gofeed.Feed, err error) *feed.Feed {
	newF := &feed.Feed{
		Url:         str.Or(f.FeedLink, url),
		Link:        f.Link,
		Title:       str.Or(f.Title, "Untitled: "+url),
		Description: f.Description,
		Error:       "",
	}
	if err != nil {
		newF.Error = err.Error()
	}
	return newF
}

func parsePosts(items []*gofeed.Item, feedTitle string) []*post.Post {
	ps := make([]*post.Post, len(items))
	for i, item := range items {
		guid := item.GUID
		if guid == "" {
			guid = hash.Sha1(item.Title + item.Description)
		}
		author := ""
		if item.Author != nil {
			author = item.Author.Name
		}
		content := item.Content
		if c := extractEncodedContent(item); len(c) > len(content) {
			content = c
		}
		var published time.Time
		if item.Published == "" && item.Updated == "" {
			published = time.Now()
		} else if item.Published != "" {
			published = *item.PublishedParsed
		} else {
			published = *item.UpdatedParsed
		}
		ps[i] = &post.Post{
			Guid:        guid,
			Author:      author,
			Link:        item.Link,
			Title:       item.Title,
			FeedTitle:   feedTitle,
			Description: item.Description,
			Content:     content,
			Published:   published,
			CreatedAt:   time.Now(),
		}
	}
	return ps
}

func extractEncodedContent(item *gofeed.Item) string {
	if item.Extensions == nil {
		return ""
	}
	if item.Extensions["content"] == nil {
		return ""
	}
	if item.Extensions["content"]["encoded"] == nil {
		return ""
	}
	for _, ext := range item.Extensions["content"]["encoded"] {
		if ext.Name == "encoded" && len(ext.Value) > 0 {
			return ext.Value
		}
	}
	return ""
}
