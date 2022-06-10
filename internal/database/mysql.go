package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type dataBaseConfig struct {
	Driver   string
	User     string
	Pass     string
	Database string
	Port     string
}

func NewMySql() DataBase {
	//TODO take values from .env
	return &dataBaseConfig{
		Driver:   "mysql",
		User:     "user",
		Pass:     "user",
		Database: "oauth_db?multiStatements=true",
		Port:     "@tcp(127.0.0.1:3306)/",
	}
}

func (config dataBaseConfig) OpenConnection() *sql.DB {
	connection, err := sql.Open(config.Driver, config.User+":"+config.Pass+config.Port+config.Database)
	if err != nil {
		panic(error.Error(err))
	}
	//doMigrations(connection)

	return connection
}

/*
func doMigrations(connection *sql.DB) {
	fileSource, err := (&file.File{}).Open("file://migrations")
	fmt.Printf("opening file error: %v \n", fileSource.Close())
	if err != nil {
		fmt.Printf("opening file error: %v \n", err)
	}

	driver, _ := mysql.WithInstance(connection, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"oauth_db",
		driver,
	)

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
*/
