package model

import "time"

type Log struct {
	ID uint `gorm:"primaryKey`
	Time time.Time
	ShortUrl string `gorm:"index"`
	Referer string
	UserAgent string
	IP string
	Country string
}
