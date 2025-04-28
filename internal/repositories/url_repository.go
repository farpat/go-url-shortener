package repositories

import (
	"database/sql"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	_ "github.com/mattn/go-sqlite3" // pour le driver SQLite
)

func List() ([]models.Url, error) {
	dbPath := config.Databases["main"]

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var urls []models.Url
	rows, err := db.Query("SELECT slug, url, created_at FROM urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url models.Url
		if err := rows.Scan(&url.Slug, &url.Url, &url.CreatedAt); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
