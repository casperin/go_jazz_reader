package urls

import "github.com/casperin/jazz_reader/internal/db"

func FindAll() ([]*Url, error) {
	urls := []*Url{}
	err := db.Conn.Select(&urls, "select * from urls")
	return urls, err
}
