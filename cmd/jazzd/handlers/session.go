package handlers

import (
	"net/http"

	"github.com/casperin/jazz_reader/internal/session"
	"github.com/spf13/viper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := session.IsLoggedIn(r)
	if loggedIn && err == nil {
		http.Redirect(w, r, "/unread", 302)
	}
	mustRenderPublicTpl("index", w, nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") != viper.GetString("password") {
		mustRenderPublicTpl("index", w, Content{"error": "Wrong password"})
		return
	}
	err := session.Login(w, r)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	http.Redirect(w, r, "/unread", 302)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	err := session.Logout(w, r)
	if err != nil {
		mustRenderErrorTpl(w, 500, err)
		return
	}
	http.Redirect(w, r, "/", 302)
}
