package handlers

import (
	"net/http"

	"github.com/casperin/jazz_reader/internal/feed"
)

func Feeds(w http.ResponseWriter, r *http.Request) {
	fs, err := feed.FindSummaries()
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	mustRenderLoggedInTpl("feeds", w, Content{
		"feeds": fs,
		"page":  "feeds",
		"title": "Feeds",
	})
}
