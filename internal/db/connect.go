package db

import (
	"fmt"

	_ "github.com/casperin/jazz_reader/internal/configuration"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var Conn *sqlx.DB

func init() {
	Conn = sqlx.MustConnect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable",
			viper.GetString("db_user"),
			viper.GetString("db_pass"),
			viper.GetString("db_name"),
		),
	)
}
