package models

import (
	"database/sql"
)

type DBEnv struct {
	User   string `envconfig:"DB_USER"`
	DbName string `envconfig:"DB_NAME"`
	Pass   string `envconfig:"DB_PASS"`
}

var Db *sql.DB

func init() {
}
