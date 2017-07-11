package handlers

import (
	"fmt"
	"net/http"

	"github.com/casperin/jazz_reader/internal/post"
)

func Saved(w http.ResponseWriter, r *http.Request) {
	ps, err := post.FindSaved()
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	mustRenderLoggedInTpl("saved", w, Content{
		"posts": ps,
		"page":  "saved",
		"title": fmt.Sprintf("%v saved", len(ps)),
	})
}
