package handlers

import (
	"net/http"
	"strconv"

	"github.com/casperin/jazz_reader/internal/feed"
	"github.com/casperin/jazz_reader/internal/post"
	"github.com/casperin/jazz_reader/internal/rss"
)

func FeedsPreview(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	f, ps, err := rss.FetchUrl(url)
	if err != nil {
		mustRenderLoggedInTpl("feeds", w, Content{
			"url":        url,
			"fetchError": err,
		})
		return
	}
	mustRenderLoggedInTpl("feeds-preview", w, Content{
		"url":   url,
		"feed":  f,
		"posts": ps,
	})
}

func FeedsAdd(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	f, ps, err := rss.FetchUrl(url)
	if err != nil {
		mustRenderLoggedInTpl("feeds-preview", w, Content{
			"error": err,
			"url":   url,
			"feed":  f,
			"posts": ps,
		})
		return
	}
	feedId, err := feed.Insert(f)
	if err != nil {
		mustRenderLoggedInTpl("feeds-preview", w, Content{
			"error": err,
			"url":   url,
			"feed":  f,
			"posts": ps,
		})
		return
	}
	for _, p := range ps {
		p.FeedId = feedId
	}
	err = post.Insert(ps)
	if err != nil {
		mustRenderLoggedInTpl("feeds-preview", w, Content{
			"error": err,
			"url":   url,
			"feed":  f,
			"posts": ps,
		})
		return
	}
	http.Redirect(w, r, "/unread", 302)
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		mustRenderErrorTpl(w, 400, err)
		return
	}
	err = post.RemoveBelongingToFeed(id)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	err = feed.RemoveById(id)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	http.Redirect(w, r, "/feeds", 302)
}
