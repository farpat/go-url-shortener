package main

import (
	"database/sql"
	"os"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var shouldSeedData bool

var rootCmd = &cobra.Command{
	Use:   "refresh-database",
	Short: "Refresh the database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := refreshDatabase(shouldSeedData); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&shouldSeedData, "seed", "s", false, "Seed the database with sample data")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func refreshDatabase(shouldSeed bool) error {
	dbPath := config.Databases["main"]

	color.Cyan("🟥  Deleting database if exists... ")
	if err := deleteDatabaseIfExists(dbPath); err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	color.Cyan("🗃️  Creating urls table...")
	if err := createUrlsTable(db); err != nil {
		return err
	}

	if shouldSeed {
		color.Cyan("🌱  Seeding database...")
		urls := []models.UrlShowItem{
			{Slug: "google", Url: "https://www.google.com"},
			{Slug: "github", Url: "https://www.github.com"},
		}

		for _, url := range urls {
			if err := storeUrl(db, url); err != nil {
				return err
			}
		}
	}

	color.Green("✅  Database refreshed successfully!")
	return nil
}

func storeUrl(db *sql.DB, url models.UrlShowItem) error {
	_, err := db.Exec("INSERT INTO urls (slug, url) VALUES (?, ?)", url.Slug, url.Url)
	return err
}

func createUrlsTable(db *sql.DB) error {
	return framework.CreateTable(db, framework.DatabaseTable{
		Name: "urls",
		Fields: []framework.DatabaseField{
			{Name: "slug", Type: "TEXT", Extra: "PRIMARY KEY"},
			{Name: "url", Type: "TEXT", Extra: "NOT NULL"},
			{Name: "created_at", Type: "TIMESTAMP", Extra: "DEFAULT CURRENT_TIMESTAMP"},
		},
	})
}

func deleteDatabaseIfExists(dbPath string) error {
	if _, err := os.Stat(dbPath); err == nil {
		return os.Remove(dbPath)
	}
	return nil
}
