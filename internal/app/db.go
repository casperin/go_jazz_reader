package app

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/casperin/jazz_reader/internal/db"
)

var ErrNotFound = errors.New("App not found")

func Get() (*App, error) {
	a := App{}
	_, err := &a, db.Conn.Get(&a, `select * from app`)
	if err == sql.ErrNoRows {
		return &a, ErrNotFound
	}
	return &a, err
}

func UpdateLastItemUpdate(appId int, t time.Time) error {
	_, err := db.Conn.Exec(
		`update app value set last_items_update=$1 where id=$2`,
		t,
		appId,
	)
	return err
}

// This just makes sure we have an app. We use this app in the db to store when
// we synced the rss' last time.
func init() {
	_, err := Get()
	if err == ErrNotFound {
		longAgo := time.Now().Add(-100 * time.Hour)
		_, err := db.Conn.Exec(
			`insert into app (last_items_update, last_email_update) values ($1, $2)`,
			longAgo, longAgo)
		if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	}
}
