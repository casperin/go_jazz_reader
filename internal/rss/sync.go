package rss

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/casperin/jazz_reader/internal/app"
	"github.com/casperin/jazz_reader/internal/feed"
	"github.com/casperin/jazz_reader/internal/post"
	"github.com/spf13/viper"
)

func SyncDaemon() error {
	for {
		a, err := app.Get()
		if err != nil {
			return fmt.Errorf("Could not get app settings:\n%v", err)
		}
		everyXMinute := viper.GetInt("sync_rss_every_x_minute")
		if everyXMinute < 10 {
			return errors.New(fmt.Sprintf("Needs a bigger delay between every syncing: %v", everyXMinute))
		}
		if time.Since(a.LastItemsUpdate) > time.Duration(everyXMinute)*time.Minute {
			err := SyncAll()
			if err != nil {
				log.Println("Error syncing:\n", err)
			}
			err = app.UpdateLastItemUpdate(a.Id, time.Now().UTC())
			if err != nil {
				log.Println("Error updating time for app:\n", err)
			}
		}
		time.Sleep(10 * time.Minute)
	}
}

func SyncAll() error {
	fs, err := feed.FindSummaries()
	if err != nil {
		return fmt.Errorf("Could not find feed summaries:\n%v", err)
	}
	var wg sync.WaitGroup
	for _, f := range fs {
		go func(feedId int, url string) {
			wg.Add(1)
			defer wg.Done()
			_, ps, err := FetchUrl(url)
			if err != nil {
				log.Println(err)
				return
			}
			for _, p := range ps {
				p.FeedId = feedId
			}
			if err := post.Insert(ps); err != nil {
				log.Println(err)
				return
			}
		}(f.Id, f.Url)
	}
	wg.Wait()
	return nil
}
