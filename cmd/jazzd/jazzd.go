package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/casperin/jazz_reader/cmd/jazzd/handlers"
	"github.com/casperin/jazz_reader/cmd/jazzd/middleware"
	"github.com/casperin/jazz_reader/internal/rss"
	"github.com/go-chi/chi"
	"github.com/gorilla/context"
	"github.com/spf13/viper"
)

func main() {
	go func() { log.Println(rss.SyncDaemon()) }()

	r := chi.NewRouter()

	r.Get("/", handlers.Index)
	r.Post("/login", handlers.LoginPost)

	// CSS
	wd, _ := os.Getwd()
	cssPath := filepath.Join(wd, "cmd", "jazzd", "static", "styles.css")
	r.Get("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cssPath)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.MustBeLoggedIn)
		r.Get("/logout", handlers.Logout)
		r.Get("/unread", handlers.Unread)
		r.Get("/feeds", handlers.Feeds)
		r.Get("/saved", handlers.Saved)
		r.Get("/feeds/{id}", handlers.Feed)
		r.Get("/read/{id}", handlers.Read)

		// feeds post
		r.Post("/preview-feed", handlers.FeedsPreview)
		r.Post("/add-feed", handlers.FeedsAdd)
		r.Post("/unsubscribe", handlers.Unsubscribe)

		// posts util
		r.Post("/mark-as-read", handlers.MarkAsRead)
		r.Get("/save/{id}", handlers.SavePost)
		r.Get("/forget/{id}", handlers.ForgetPost)
		r.Post("/save-url", handlers.SaveUrl)
		r.Get("/forget-url/{id}", handlers.ForgetUrl)
	})

	fmt.Printf("Listening on port %v\n", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), context.ClearHandler(r)))
}
