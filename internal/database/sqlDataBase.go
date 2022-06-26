package database

import (
	"database/sql"
	"errors"
	"gaston.frassatti/aouth_manager/internal/handlers"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Config struct {
	Driver   string
	User     string
	Pass     string
	Database string
	Port     string
}

type mySql struct {
	sql *sql.DB
}

func NewConnection(config Config) *mySql {
	db, err := sql.Open(config.Driver, config.User+":"+config.Pass+config.Port+config.Database)
	if err != nil {
		panic(error.Error(err))
	}
	//doMigrations(connection)

	return &mySql{sql: db}
}

func (db mySql) GetGrants(uuid string) (grants handlers.Grant, err error) {
	row := db.sql.QueryRow("SELECT * FROM oauth_db.grants WHERE oauth_uuid = ?", uuid)

	err = row.Scan(&grants.Uuid, &grants.AccessToken, &grants.ExpiresDate, &grants.TokenType)
	if errors.Is(err, sql.ErrNoRows) {
		return grants, nil
	}
	return grants, err
}

func (db mySql) Upsert(grants handlers.Grant) {
	stmt, err := db.sql.Prepare("INSERT INTO oauth_db.grants (oauth_uuid, acces_token, expires_date, token_type) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE acces_token=values( acces_token), expires_date=values(expires_date)")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(grants.Uuid, grants.AccessToken, grants.ExpiresDate, grants.TokenType)
}
