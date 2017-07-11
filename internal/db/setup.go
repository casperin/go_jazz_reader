package db

func Create() {
	Conn.MustExec(appTable)
	Conn.MustExec(feedTable)
	Conn.MustExec(itemTable)
}

func Drop() {
	Conn.Exec(dropItemTable)
	Conn.Exec(dropFeedTable)
	Conn.Exec(dropAppTable)
}

func Reset() {
	Drop()
	Create()
}
