package handlers

import (
	"fmt"
	"net/http"

	"github.com/casperin/jazz_reader/internal/post"
	"github.com/casperin/jazz_reader/internal/urls"
)

func Saved(w http.ResponseWriter, r *http.Request) {
	ps, err := post.FindSaved()
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	us, err := urls.FindAll()
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	mustRenderLoggedInTpl("saved", w, Content{
		"posts": ps,
		"urls":  us,
		"page":  "saved",
		"title": fmt.Sprintf("%v saved", len(ps)+len(us)),
	})
}
