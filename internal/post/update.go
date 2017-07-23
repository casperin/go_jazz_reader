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
	if err != nil {
		return fmt.Errorf("Could mark items as read:\n%v", err)
	}
	return nil
}

func SetSaved(id int, saved bool) error {
	sql := `update posts set saved=$1 where id=$2`
	_, err := db.Conn.Exec(sql, saved, id)
	if err != nil {
		return fmt.Errorf("Could not set saved:\n%v", err)
	}
	return nil
}
