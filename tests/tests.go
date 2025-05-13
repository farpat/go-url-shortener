package tests

import (
	"database/sql"
	"os"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
	_ "github.com/mattn/go-sqlite3"
)

func SetupTestDB() (teardown func(), db *sql.DB) {
	dbPath := "database_test.db"
	config.Databases["main"] = dbPath
	absoluteDbPath := framework.ProjectPath(dbPath)

	db, err := sql.Open("sqlite3", absoluteDbPath)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		slug TEXT PRIMARY KEY,
		url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	return func() {
		db.Close()
		os.Remove(absoluteDbPath)
	}, db
}
