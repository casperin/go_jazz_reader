package db

import (
	"fmt"
	"os"

	_ "github.com/casperin/jazz_reader/internal/configuration"
	"github.com/casperin/jazz_reader/internal/util/str"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	base       = "user=%s password=%s dbname=%s sslmode=disable"
	dbUser     = viper.GetString("db_user")
	dbPass     = viper.GetString("db_pass")
	dbName     = str.Or(viper.GetString("db_name"), "jazzreader_dev")
	testDbUser = viper.GetString("test_db_user")
	testDbPass = viper.GetString("test_db_pass")
	testDbName = str.Or(viper.GetString("test_db_name"), "jazzreader_test")
	Conn       *sqlx.DB
)

func init() {
	connect()
}

func connect() {
	var cs string
	if os.Getenv("ENV") == "TEST" {
		cs = fmt.Sprintf(base, testDbUser, testDbPass, testDbName)
	} else {
		cs = fmt.Sprintf(base, dbUser, dbPass, dbName)
	}
	Conn = sqlx.MustConnect("postgres", cs)
}
