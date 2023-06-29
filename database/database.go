package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitializeTGSCADatabase(dbname, username, password, dbhost, dbport string) (*sqlx.DB, error) {
	conn := username + ":" + password + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	tgscadb, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, fmt.Errorf("Error in initializing cad database: %s", err)
	}

	return tgscadb, nil
}
