package dao

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ushorten/ushorten/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"strings"
)

func NewMySQL(conf *model.Config) (*gorm.DB, error) {
	var credential, protocol, address string
	address = conf.DbHost
	if conf.DbPort != 3306 {
		address += ":" + string(conf.DbPort)
	}
	protocol = "tcp"
	if strings.Contains(address, "/") {
		protocol = "unix"
	}
	credential = conf.DbUser
	if conf.DbPasswd != "" {
		credential += ":" + conf.DbPasswd
	}
	dsn := fmt.Sprintf(
		"%s@%s(%s)/%s?collation=%s&loc=local&parseTime=true",
		credential,
		protocol,
		address,
		conf.DbName,
		conf.DbCollation,
	)
	logrus.Debug("dsn", dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
