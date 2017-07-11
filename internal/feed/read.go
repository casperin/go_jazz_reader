package feed

import (
	"database/sql"
	"errors"

	"github.com/casperin/jazz_reader/internal/db"
)

const selectWhereBase = "select * from feeds where "

var ErrNotFound = errors.New("Feed not found")

func FindSummaries() ([]*Feed, error) {
	fs := []*Feed{}
	err := db.Conn.Select(&fs, "select id, title, url from feeds order by id desc")
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return fs, err
}

func FindById(id int) (*Feed, error) {
	f := Feed{}
	err := db.Conn.Get(&f, selectWhereBase+"id=$1", id)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return &f, err
}

func FindByUrl(url string) (*Feed, error) {
	f := Feed{}
	err := db.Conn.Get(&f, selectWhereBase+"url=$1", url)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return &f, err
}
