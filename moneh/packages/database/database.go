package database

import (
	"database/sql"
	"moneh/configs"

	_ "github.com/go-sql-driver/mysql"
)

var dbs *sql.DB
var err error

func Init() {
	conf := configs.GetConfig()

	// username:password@protocol(address)/dbname
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	dbs, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connection error!!")
	}

	err = dbs.Ping()
	if err != nil {
		panic("DSN Invalid!!")
	}
}

func CreateCon() *sql.DB {
	return dbs
}
