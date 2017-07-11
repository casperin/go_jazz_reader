package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var store = sessions.NewCookieStore([]byte(viper.GetString("app_secret")))

func Login(w http.ResponseWriter, r *http.Request) error {
	// See: https://github.com/gorilla/sessions/issues/16 why error is ignored
	us, _ := store.Get(r, "logged-in")
	us.Values["logged-in"] = true
	return us.Save(r, w)
}

func IsLoggedIn(r *http.Request) (bool, error) {
	us, err := store.Get(r, "logged-in")
	if err != nil {
		return false, err
	}
	return us.Values["logged-in"] == true, nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	us, err := store.Get(r, "logged-in")
	if err != nil {
		return err
	}
	us.Values["logged-in"] = nil
	return us.Save(r, w)
}
