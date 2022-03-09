package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type DbEnv struct {
	User string
	Name string
	Pass string
}

var Db *sqlx.DB

func init() {
	var dbEnv DbEnv
	envconfig.Process("db", &dbEnv)
	// env: DB_USER -> dbEnv.User
	ds := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbEnv.User, dbEnv.Name, dbEnv.Pass)
	var err error
	Db, err = sqlx.Open("postgres", ds)
	if err != nil {
		panic(err)
	}
}
