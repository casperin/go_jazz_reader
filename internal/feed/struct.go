package feed

type Feed struct {
	Id          int    `db:"id" json:"id"`
	Url         string `db:"url" json:"url"`
	Link        string `db:"link" json:"link"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Error       string `db:"error" json:"error"`
}
