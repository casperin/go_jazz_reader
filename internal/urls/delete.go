package urls

import "github.com/casperin/jazz_reader/internal/db"

func Remove(id int) error {
	sql := `delete from urls where id=$1`
	_, err := db.Conn.Exec(sql, id)
	return err
}
