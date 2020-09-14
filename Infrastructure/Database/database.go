package Database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	username     string
	password     string
	name         string
	databaseType string
}

func (database *Database) formatSqlConnection() string {
	return database.username + ":" + database.password + "@/" + database.name

}

func (database *Database) GetConnection() *sql.DB {
	db, err := sql.Open(database.databaseType, database.formatSqlConnection())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func NewDatabase() *Database {
	return &Database{
		username:     "root",
		password:     "123",
		name:         "hostgator",
		databaseType: "mysql",
	}
}
