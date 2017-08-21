package urls

import "time"

type Url struct {
	Id        int       `db:"id" json:"id"`
	Url       string    `db:"url" json:"url"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
