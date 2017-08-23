package urls

import (
	"time"

	"github.com/casperin/jazz_reader/internal/db"
)

func Save(u, title string) error {
	sql := `insert into urls (url, title, created_at) values ($1, $2, $3)`
	_, err := db.Conn.Exec(sql, u, title, time.Now())
	return err
}
