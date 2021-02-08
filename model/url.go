package model

import "time"

type Url struct {
	Keyword string `gorm:"primaryKey"`
	Url string
	Title string
	IP string `gorm:"index"`
	Clicks uint64
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
