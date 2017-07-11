package middleware

import (
	"net/http"

	"github.com/casperin/jazz_reader/internal/session"
)

func MustBeLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok, err := session.IsLoggedIn(r); !ok || err != nil {
			http.Redirect(w, r, "/", 302)
			return
		}
		next.ServeHTTP(w, r)
	})
}
