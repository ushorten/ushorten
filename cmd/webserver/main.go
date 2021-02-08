package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ushorten/ushorten/autoload"
	"github.com/ushorten/ushorten/model"
	"github.com/ushorten/ushorten/service/dao"
	"os"
)

var config *model.Config

func init() {
	var err error
	config = autoload.Config
	config.AppPath, err = os.Getwd()
	if err != nil {
		logrus.Panic(err)
	}
	dao.DB, err = dao.NewDB(config)
	if err != nil {
		logrus.Panic(err)
	}
	if config.AppDebug {
		dao.DB = dao.DB.Debug()
	}

	startup()
}

func main() {
	logrus.Info(os.Getenv("PWD"))
	logrus.Infof("%+v", config)
}

func startup() {
	dao.DB.AutoMigrate(
		model.Url{},
		model.Option{},
		model.Log{},
	)
}
