package database

import (
	"database/sql"
)

type DataBase interface {
	OpenConnection() *sql.DB
}
