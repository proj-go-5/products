package storage

import (
	"fmt"
	"os"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLStorage struct {
	db *sqlx.DB
}

func NewStorage() *MySQLStorage {
	s := &MySQLStorage{}

	dbPort, ok := os.LookupEnv("MYSQL_TCP_PORT_EXPOSE")
	if !ok {
		fmt.Println("env var MYSQL_TCP_PORT_EXPOSE is not found")
		dbPort = "3306"
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		fmt.Println("env var DB_HOST is not found")
		dbHost = "localhost"
	}

	dbConf := mysqlDriver.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName:               "products",
		AllowNativePasswords: true,
	}

	s.db = DBConnect(dbConf)
	err := setUpToDateDB(s.db)
	if err != nil {
		fmt.Println(err)
	}
	return s
}

/*
example usage :

products := []Product{}

// receive all products
storage.Get(products, "Product", "")

// receive all products with id=3
storage.Get(products, "Product", "id=3")

// receive all products with filter
storage.Get(products, "Product", "title LIKE '%tro%'")
*/
func (ms *MySQLStorage) Get(dest interface{}, tableName string, filter string) error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	err := get(ms.db, dest, tableName, filter)
	return err
}

func (ms *MySQLStorage) Close() error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	return ms.db.Close()
}
