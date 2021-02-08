package model

type Option struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"index"`
	Value string
}
