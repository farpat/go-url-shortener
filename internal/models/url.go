package models

import "time"

type Url struct {
	Slug      string
	Url       string
	CreatedAt time.Time
}
