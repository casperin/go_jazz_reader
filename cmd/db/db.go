package main

import (
	"fmt"
	"os"

	"github.com/casperin/jazz_reader/internal/db"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("I need a command: create, drop, reset")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "create":
		db.Create()
		fmt.Println("Database created succesfully")
	case "drop":
		db.Drop()
		fmt.Println("Database dropped")
	case "reset":
		db.Reset()
		fmt.Println("Database reset succesfully")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
