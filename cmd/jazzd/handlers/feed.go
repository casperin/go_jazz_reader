package handlers

import (
	"net/http"
	"strconv"

	"github.com/casperin/jazz_reader/internal/feed"
	"github.com/casperin/jazz_reader/internal/post"
	"github.com/go-chi/chi"
)

func Feed(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	f, err := feed.FindById(id)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	ps, err := post.FindByFeedId(id)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	mustRenderLoggedInTpl("feed", w, Content{
		"feed":  f,
		"posts": ps,
		"title": f.Title,
	})
}
