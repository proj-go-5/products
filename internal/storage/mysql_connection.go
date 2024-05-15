package storage

import (
	"fmt"
	"strings"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/proj-go-5/products/internal/dto"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func DBConnect(dbConf mysqlDriver.Config) *sqlx.DB {
	db, err := sqlx.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func get(db *sqlx.DB, dest interface{}, tableName string, filter string, args ...interface{}) error {
	query := "SELECT * FROM " + tableName
	if filter != "" {
		query += " WHERE " + filter
	}
	err := db.Select(dest, query, args...)
	return err
}

func add(db *sqlx.DB, values map[string]interface{}, tableName string) error {
	// Constructing a sql request
	//like this : INSERT INTO Product (first_name,last_name,email) VALUES (:first,:last,:email)
	//values := map[string]interface{}{
	//		"first": "Bin",
	//		"last":  "Smuth",
	//		"email": "bensmith@allblacks.nz",
	//	}
	sqlFields := make([]string, 0, len(values))
	sqlValues := make([]string, 0, len(values))
	for k := range values {
		sqlFields = append(sqlFields, k)
		sqlValues = append(sqlValues, fmt.Sprintf(":%v", k))
	}

	request := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(sqlFields, ","),
		strings.Join(sqlValues, ","),
	)

	_, err := db.NamedExec(request, values)
	return err
}

func updateProduct(db *sqlx.DB, product *dto.ProductRequest) error {
	_, err := db.NamedExec(`UPDATE Product SET title=:title, price=:price, description=:description, update_date=CURRENT_TIMESTAMP(), images=:image WHERE id=:id`, product)
	return err
}

func delete(db *sqlx.DB, tableName string, id int32) error {
	_, err := db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE id=%d`, tableName, id))
	return err
}

func setUpToDateDB(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("cannot obtain driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://././migrations",
		"products", driver)
	if err != nil {
		return fmt.Errorf("cannot migrate: %s", err)
	}
	return m.Up()
}

func (ms *MySQLStorage) GetProducts(filter string, sortBy string, order string, pageNum int, pageSize int, ids string) ([]dto.ProductRequest, error) {
	var products []dto.ProductRequest
	queryArgs := []interface{}{}

	searchParam := ""
	if ids != "" {
		idsSlice := strings.Split(ids, ",")
		for _, id := range idsSlice {
			queryArgs = append(queryArgs, id)
		}
		searchParam = fmt.Sprintf("WHERE id IN (?%s)", strings.Repeat(",?", len(idsSlice)-1))

	} else if filter != "" {
		searchParam = "WHERE title LIKE ? OR description LIKE ?"
		queryArgs = append(queryArgs, "%"+filter+"%", "%"+filter+"%")
	}

	sortBySQL := "ORDER BY name"
	if sortBy != "" {
		sortBySQL = "ORDER BY " + sortBy
	}

	orderSQL := "ASC"
	if order == "desc" {
		orderSQL = "DESC"
	}

	offset := (pageNum - 1) * pageSize
	SQLRequest := fmt.Sprintf("SELECT * FROM Product %s %s %s LIMIT ? OFFSET ?", searchParam, sortBySQL, orderSQL)
	queryArgs = append(queryArgs, pageSize, offset)

	err := ms.db.Select(&products, SQLRequest, queryArgs...)
	if err != nil {
		return nil, err
	}

	return products, nil
}
