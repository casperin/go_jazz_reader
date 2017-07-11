package post

import (
	"database/sql"
	"errors"

	"github.com/casperin/jazz_reader/internal/db"
)

const (
	selectWhereBase = "select * from posts where "
	orderDesc       = " order by id desc"
)

var ErrNotFound = errors.New("Post not found")

func findBySql(sqlStr string, value interface{}) ([]*Post, error) {
	p := []*Post{}
	err := db.Conn.Select(&p, sqlStr, value)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return p, err
}

func FindByFeedId(feedId int) ([]*Post, error) {
	return findBySql(selectWhereBase+"feed_id=$1"+orderDesc, feedId)
}

func FindByStatus(status string) ([]*Post, error) {
	return findBySql(selectWhereBase+"status=$1"+orderDesc, status)
}

func FindSaved() ([]*Post, error) {
	return findBySql(selectWhereBase+"saved=$1"+orderDesc, true)
}

// Single post
func FindById(id string) (*Post, error) {
	p := Post{}
	err := db.Conn.Get(&p, selectWhereBase+"id=$1", id)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return &p, err
}
