package drivers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
// Open a connection to the database 
func OpenDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
