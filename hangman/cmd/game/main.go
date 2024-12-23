package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


const (
	schemaSQL = `
	CREATE TABLE IF NOT EXISTS words (
		id INTEGER,
		word VARCHAR(100),
		category VARCHAR(50)
	)
	`
	insertSQL = `
	INSERT INTO words (id, word, category) VALUES (?, ?, ?, ?)`
)

var tmpl *tempalte.Template


type DB struct {
	sql *sql.DB
	stmt *sql.Stmt
	buffer []Word
}

type Word struct {
	ID   int
	Word string
	Category string
}

func NewDB(dbfile string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	if _, err := sqlDB.Exec(schemaSQL); err != nil {
		return nil, err
	}
	stmt, err := sqlDB.Prepare(insertSQL)
		if err != nil {
			return nil, err
		}
	db := DB{
		sql: sqlDB,
		stmt: stmt,
		buffer: make([]Word,0, 1024)
	}

	return &db
}


func main() {


}
