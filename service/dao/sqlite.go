package dao

import (
	"fmt"
	"github.com/ushorten/ushorten/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

func NewSqlite(conf *model.Config) (*gorm.DB, error) {
	path := filepath.ToSlash(conf.DbPath)
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("%s%s%s", conf.AppPath, string(os.PathSeparator), filepath.FromSlash(conf.DbPath))
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), 0644)
	}
	dsn := path
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
