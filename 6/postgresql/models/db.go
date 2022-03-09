package models

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type DbEnv struct {
	User string
	Name string
	Pass string
}

var Db *sql.DB

func init() {
	var dbEnv DbEnv
	envconfig.Process("db", &dbEnv)
	ds := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbEnv.User, dbEnv.Name, dbEnv.Pass)
	var err error
	Db, err = sql.Open("postgres", ds)
	if err != nil {
		panic(err)
	}
}
