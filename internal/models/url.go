package models

import "time"

type Url struct {
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
