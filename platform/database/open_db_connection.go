package database

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func OpenDBConnection() (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mysql":
		db, err = MysqlConnection()
	case "postgresql":
		db, err = PostgreSQLConnection()
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
