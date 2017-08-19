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
	ids := ""
	for i, p := range ps {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(p.Id)
	}
	mustRenderLoggedInTpl("/unread", w, Content{
		"count": len(ps),
		"ids":   ids,
		"page":  "unread",
		"posts": ps,
		"title": fmt.Sprintf("%v unread", len(ps)),
	})
}
