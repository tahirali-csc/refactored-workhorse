package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/workhorse/apiserver/pkg/config"
)

type RunFunc func(db *sql.DB) error

func Run(runFunc RunFunc) error{
	appConfig := config.GetAppConfig()
	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		appConfig.Database.Name, appConfig.Database.User, appConfig.Database.Password, appConfig.Database.Host)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	return runFunc(db)
}
