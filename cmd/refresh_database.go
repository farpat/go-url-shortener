package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/services"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3" // pour le driver SQLite
)

func main() {
	dbPath := config.Databases["main"]

	color.Cyan("🟥 Deleting database if exists... ")
	if err := deleteDatabaseIfExists(dbPath); err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	color.Cyan("🗃️  Creating urls table...")
	if err := createUrlsTable(db); err != nil {
		panic(err)
	}

	if !shouldSeedDatabase() {
		return
	}

	color.Cyan("🌱 Seeding database...")

	for _, url := range []string{
		"https://golang.org",
		"https://www.google.com",
		"https://github.com",
	} {
		slug := services.GenerateSlug(url)

		url := models.Url{
			Slug:      slug,
			Url:       url,
			CreatedAt: time.Now(),
		}

		if err := storeUrl(db, url); err != nil {
			panic(err)
		}
	}

	color.Green("🪄  Database refreshed successfully.")
}

func shouldSeedDatabase() bool {
	return os.Getenv("SEED") == "true" || os.Getenv("SEED") == "1"
}

func storeUrl(db *sql.DB, url models.Url) error {
	_, err := db.Exec("INSERT INTO urls (slug, url, created_at) VALUES (?, ?, ?)", url.Slug, url.Url, url.CreatedAt)
	return err
}

func createUrlsTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS urls (
			slug TEXT PRIMARY KEY,
			url TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := db.Exec(createTableQuery)
	return err
}

func deleteDatabaseIfExists(dbPath string) error {
	if _, err := os.Stat(dbPath); err == nil {
		return os.Remove(dbPath)
	}
	return nil
}
