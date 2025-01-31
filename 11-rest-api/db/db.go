package db

/**
go-sqlite3 is a SQL driver that we needed to have in the project to use an SQLite database.
We will not directly interact with the go-sqlite3 package. We will directly interact with
database/sql package (which is a part of the Go standard library)
and Go will establish the connection between the database/sql package and go-sqlite3 package.

Since the go-sqlite3 package will not be mentioned in this file, to prevent removing the package
from this file the underscore (_) in front of the github.com/mattn/go-sqlite3 package was used.
*/

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// sqlite3 - means that we will use this driver
	// api.db - name of the database, which will be created automatically if does not exist
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
