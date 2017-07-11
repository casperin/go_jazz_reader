package feed

import "github.com/casperin/jazz_reader/internal/db"

func RemoveById(id int) error {
	sql := `delete from feeds where id=$1`
	_, err := db.Conn.Exec(sql, id)
	return err
}
