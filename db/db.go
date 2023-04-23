package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"occupie/config"
)

var occupie *sql.DB
var err error

func NewDbConnection(env string) {
	appConfig := config.GlobalConfig(env)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		appConfig.Db.Host, appConfig.Db.Port, appConfig.Db.User, appConfig.Db.Password, appConfig.Db.Name)

	occupie, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Load DB successfully!!")
	defer occupie.Close()
}

func OccupieDB() *sql.DB {
	return occupie
}
