package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBConnect() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:13306)/products")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return db
}
