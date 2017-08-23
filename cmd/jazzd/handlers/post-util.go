package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/casperin/jazz_reader/internal/post"
	"github.com/casperin/jazz_reader/internal/util/str"
	"github.com/go-chi/chi"
)

func SavePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	err = post.SetSaved(id, true)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	to := str.Or(r.FormValue("to"), "read/"+idStr)
	http.Redirect(w, r, "/"+to, 302)
}

func ForgetPost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	err = post.SetSaved(id, false)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	to := str.Or(r.FormValue("to"), "read/"+idStr)
	http.Redirect(w, r, "/"+to, 302)
}

func MarkAsRead(w http.ResponseWriter, r *http.Request) {
	idsString := r.FormValue("ids")
	idsStrList := strings.Split(idsString, ",")
	ids := make([]int, len(idsStrList))
	for i, id := range idsStrList {
		intId, _ := strconv.Atoi(id)
		ids[i] = intId
	}
	if err := post.MarkAsRead(ids); err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	http.Redirect(w, r, "/unread", 302)
}
