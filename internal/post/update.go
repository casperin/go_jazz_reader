package post

import (
	"fmt"

	"github.com/casperin/jazz_reader/internal/db"
)

func MarkAsRead(ids []int) error {
	args := make([]interface{}, len(ids))
	sql := `update posts set status='read' where id in (`
	for i, id := range ids {
		args[i] = id
		if i > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("$%v", i+1)
	}
	sql += ")"
	_, err := db.Conn.Exec(sql, args...)
	return err
}

func SetSaved(id int, saved bool) error {
	sql := `update posts set saved=$1 where id=$2`
	_, err := db.Conn.Exec(sql, saved, id)
	return err
}
