package storage

import "database/sql"

type MySQLStorage struct {
	db *sql.DB
}
