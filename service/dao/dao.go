package dao

import (
	"github.com/ushorten/ushorten/model"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func NewDB(conf *model.Config) (*gorm.DB, error) {
	switch conf.DbType {
	case "mysql":
		return NewMySQL(conf)
	default:
		return NewSqlite(conf)
	}
}
