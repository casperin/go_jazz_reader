package post

import (
	"fmt"

	"github.com/casperin/jazz_reader/internal/db"
)

func Insert(ps []*Post) error {
	sql := `insert into posts
	(feed_id, guid, author, link, title, feed_title, description, content, published, created_at) values `
	n := 0
	values := []interface{}{}
	for i, j := 0, len(ps)-1; i < j; i, j = i+1, j-1 {
		ps[i], ps[j] = ps[j], ps[i]
	}
	for i, p := range ps {
		if i != 0 {
			sql += ", "
		}
		sql += fmt.Sprintf(
			"($%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v)",
			n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8, n+9, n+10,
		)
		n += 10
		values = append(values, []interface{}{
			p.FeedId,
			p.Guid,
			p.Author,
			p.Link,
			p.Title,
			p.FeedTitle,
			p.Description,
			p.Content,
			p.Published,
			p.CreatedAt,
		}...)
	}
	sql += " on conflict do nothing;"
	_, err := db.Conn.Exec(sql, values...)
	return err
}
