package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/casperin/jazz_reader/internal/post"
)

func Unread(w http.ResponseWriter, r *http.Request) {
	ps, err := post.FindByStatus("new")
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	gps := map[string][]*post.Post{}
	ids := ""
	for i, p := range ps {
		if _, ok := gps[p.FeedTitle]; !ok {
			gps[p.FeedTitle] = []*post.Post{}
		}
		gps[p.FeedTitle] = append(gps[p.FeedTitle], p)
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(p.Id)
	}
	mustRenderLoggedInTpl("/unread", w, Content{
		"count": len(ps),
		"gps":   gps,
		"ids":   ids,
		"page":  "unread",
		"title": fmt.Sprintf("%v unread", len(ps)),
	})
}
