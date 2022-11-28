package bootstrap

import (
	"echain-user-updater/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func InitDb() *sqlx.DB {
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	pool, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = pool.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("connected to db")

	return pool
}
