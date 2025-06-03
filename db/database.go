package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS temperature (
        id TEXT PRIMARY KEY,
        temperature REAL,
        date TEXT,
        time TEXT,
        location TEXT
    );`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
