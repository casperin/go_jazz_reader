package app

import "time"

type App struct {
	Id              int       `db:"id"`
	LastItemsUpdate time.Time `db:"last_items_update"`
	LastEmailUpdate time.Time `db:"last_email_update"`
}
