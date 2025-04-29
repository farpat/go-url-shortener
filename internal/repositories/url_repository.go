package repositories

import (
	"database/sql"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/utils"
	_ "github.com/mattn/go-sqlite3" // pour le driver SQLite
)

type NotFoundError struct {
	Slug string
}

func (e *NotFoundError) Error() string {
	return "URL linked to '" + e.Slug + "' not found"
}

func All() ([]models.UrlListItem, error) {
	db := openDB()

	urls := []models.UrlListItem{}
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
	db := openDB()

	var url models.UrlShowItem
	err := db.QueryRow("SELECT slug, url, created_at FROM urls WHERE slug = ?", slug).Scan(&url.Slug, &url.Url, &url.CreatedAt)
	if err != nil {
		return models.UrlShowItem{}, &NotFoundError{Slug: slug}
	}

	return url, nil
}

func Delete(slug string) error {
	db := openDB()

	result, err := db.Exec("DELETE FROM urls WHERE slug = ?", slug)
	if err != nil {
		return err
	}

	rowsAffectedCount, _ := result.RowsAffected()
	if rowsAffectedCount == 0 {
		return &NotFoundError{Slug: slug}
	}

	return nil
}

func openDB() *sql.DB {
	dbPath := config.Databases["main"]

	db, err := sql.Open("sqlite3", utils.ProjectPath(dbPath))
	if err != nil {
		panic(err)
	}

	return db
}
