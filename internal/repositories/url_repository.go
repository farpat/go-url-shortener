package repositories

import (
	"database/sql"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	_ "github.com/mattn/go-sqlite3" // pour le driver SQLite
)

func All() ([]models.UrlListItem, error) {
	dbPath := config.Databases["main"]

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var urls []models.UrlListItem
	rows, err := db.Query("SELECT slug, url FROM urls ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url models.UrlListItem
		if err := rows.Scan(&url.Slug, &url.Url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func Find(slug string) (models.UrlShowItem, error) {
	dbPath := config.Databases["main"]

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return models.UrlShowItem{}, err
	}
	defer db.Close()

	var url models.UrlShowItem
	err = db.QueryRow("SELECT slug, url, created_at FROM urls WHERE slug = ?", slug).Scan(&url.Slug, &url.Url, &url.CreatedAt)
	if err != nil {
		return models.UrlShowItem{}, err
	}

	return url, nil
}
