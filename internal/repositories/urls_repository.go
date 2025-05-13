package repositories

import (
	"database/sql"
	"time"

	"github.com/farpat/go-url-shortener/internal/config"
	internalErrors "github.com/farpat/go-url-shortener/internal/errors"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/services/string_utils"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
	_ "github.com/mattn/go-sqlite3"
)

type UrlRepository struct {
}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{}
}

func (r *UrlRepository) All() ([]models.UrlListItem, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

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

func (r *UrlRepository) Exists(slug string) (bool, error) {
	db, err := openDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var count int
	db.QueryRow("SELECT COUNT(*) FROM urls WHERE slug = ?", slug).Scan(&count)
	return count > 0, nil
}

func (r *UrlRepository) Find(slug string) (models.UrlShowItem, error) {
	db, err := openDB()
	if err != nil {
		return models.UrlShowItem{}, err
	}
	defer db.Close()

	var url models.UrlShowItem
	err = db.QueryRow("SELECT slug, url, created_at FROM urls WHERE slug = ?", slug).Scan(&url.Slug, &url.Url, &url.CreatedAt)
	if err != nil {
		return models.UrlShowItem{}, &internalErrors.NotFoundError{Slug: slug}
	}

	return url, nil
}

func (r *UrlRepository) Delete(slug string) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM urls WHERE slug = ?", slug)
	if err != nil {
		return err
	}

	rowsAffectedCount, _ := result.RowsAffected()
	if rowsAffectedCount == 0 {
		return &internalErrors.NotFoundError{Slug: slug}
	}

	return nil
}

func (r *UrlRepository) Create(url models.UrlShowItem) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if url.Slug == "" {
		url.Slug = string_utils.GenerateSlug(url.Url)
	}

	if url.CreatedAt.IsZero() {
		url.CreatedAt = time.Now()
	}

	_, err = db.Exec("INSERT INTO urls (slug, url, created_at) VALUES (?, ?, ?)", url.Slug, url.Url, url.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func openDB() (*sql.DB, error) {
	dbPath := config.Databases["main"]

	db, err := sql.Open("sqlite3", framework.ProjectPath(dbPath))
	if err != nil {
		return nil, err
	}
	return db, nil
}
