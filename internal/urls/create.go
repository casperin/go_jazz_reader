package urls

import (
	"time"

	"github.com/casperin/jazz_reader/internal/db"
)

func Save(u string) error {
	sql := `insert into urls (url, created_at) values ($1, $2)`
	_, err := db.Conn.Exec(sql, u, time.Now())
	return err
}
