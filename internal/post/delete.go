package post

import "github.com/casperin/jazz_reader/internal/db"

func RemoveBelongingToFeed(feedId int) error {
	sql := `delete from posts where feed_id=$1`
	_, err := db.Conn.Exec(sql, feedId)
	return err
}
