package models

import (
	"encoding/json"
	"time"
)

type UrlShowItem struct {
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func (u UrlShowItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Slug      string `json:"slug"`
		Url       string `json:"url"`
		CreatedAt string `json:"created_at"`
	}{
		Slug:      u.Slug,
		Url:       u.Url,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type UrlListItem struct {
	Slug string `json:"slug"`
	Url  string `json:"url"`
}
