package storage

import (
	"fmt"
	"os"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/proj-go-5/products/internal/dto"
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

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		fmt.Println("env var DB_USER is not found")
		dbHost = "root"
	}
	dbPass := os.Getenv("DB_PASSWORD")

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		fmt.Printf("env var DB_NAME is not found")
		dbName = "products"
	}

	dbConf := mysqlDriver.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName:               dbName,
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

func (ms *MySQLStorage) Get(dest interface{}, tableName string, filter string, args ...interface{}) error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	err := get(ms.db, dest, tableName, filter, args...)
	return err
}

func (ms *MySQLStorage) Add(values map[string]interface{}, tableName string) error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	err := add(ms.db, values, tableName)
	return err
}

func (ms *MySQLStorage) UpdateProduct(product *dto.ProductRequest) error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	err := updateProduct(ms.db, product)
	return err
}

func (ms *MySQLStorage) Delete(tableName string, id int32) error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	err := delete(ms.db, tableName, id)
	return err
}

func (ms *MySQLStorage) Close() error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	return ms.db.Close()
}
