package post

import (
	"html/template"
	"time"
)

type Post struct {
	Id          int       `db:"id" json:"id"`
	Status      string    `db:"status" json:"status"`
	Saved       bool      `db:"saved" json:"saved"`
	FeedId      int       `db:"feed_id" json:"feed_id"`
	Guid        string    `db:"guid" json:"guid"`
	Author      string    `db:"author" json:"author"`
	Link        string    `db:"link" json:"link"`
	Title       string    `db:"title" json:"title"`
	FeedTitle   string    `db:"feed_title" json:"feed_title"`
	Description string    `db:"description" json:"description"`
	Content     string    `db:"content" json:"content"`
	Published   time.Time `db:"published" json:"published"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

func (p *Post) ContentHuman() template.HTML {
	if p.Content == "" {
		return template.HTML(p.Description)
	}
	return template.HTML(p.Content)
}

func (p *Post) PublishedHuman() string {
	return p.Published.Format("Jan 2, 2006; 15:04")
}
