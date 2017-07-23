package feed

import (
	"fmt"

	"github.com/casperin/jazz_reader/internal/db"
)

func Insert(f *Feed) (int, error) {
	found, err := FindByUrl(f.Url)
	if err != ErrNotFound {
		if found != nil {
			return found.Id, err
		}
		return 0, err
	}
	stmt, err := db.Conn.PrepareNamed(`
	INSERT INTO feeds
	(url, link, title, description, error)
	VALUES
	(:url, :link, :title, :description, :error)
	RETURNING id`)
	if err != nil {
		return 0, fmt.Errorf("Could not insert feed %s:\n%v", f.Url, err)
	}

	var id int
	err = stmt.Get(&id, map[string]interface{}{
		"url":         f.Url,
		"link":        f.Link,
		"title":       f.Title,
		"description": f.Description,
		"error":       f.Error,
	})
	if err != nil {
		return id, fmt.Errorf("Could not get Feed after %s inserting it:\n%v", f.Url, err)
	}
	return id, nil
}
