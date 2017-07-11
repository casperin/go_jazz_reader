package main

import (
	"log"

	"github.com/casperin/jazz_reader/internal/rss"
)

func main() {
	err := rss.SyncAll()

	if err != nil {
		log.Fatal(err)
	}
}
