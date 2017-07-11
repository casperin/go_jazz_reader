package handlers

import (
	"net/http"
	"net/url"

	"github.com/casperin/jazz_reader/internal/post"
	"github.com/go-chi/chi"
)

func Read(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := post.FindById(id)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	var base string
	if itemLink, err := url.Parse(p.Link); err == nil {
		base = itemLink.Scheme + "://" + itemLink.Host
	}
	mustRenderLoggedInTpl("read", w, Content{
		"title": p.Title,
		"post":  p,
		"base":  base,
	})
}
