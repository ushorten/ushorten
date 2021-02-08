package autoload

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/ushorten/ushorten/model"
	"os"
	"strconv"
)

var Log *logrus.Logger
var Config *model.Config

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout

	Config = model.NewConfig()

	err := godotenv.Load()
	if err != nil {
		Log.Info("loaded environment variables from env")
		isDebug, err := strconv.ParseBool(env("APP_DEBUG", "false"))
		if err != nil {
			isDebug = false
		}
		Config.AppDebug = isDebug
		Config.DbType = env("DB_TYPE", "sqlite")
		Config.DbPath = env("DB_PATH", Config.DbPath)
		Config.DbHost = env("DB_HOST", Config.DbHost)
		port, err := strconv.Atoi(env("DB_HOST", strconv.FormatInt(int64(Config.DbPort), 10)))
		if err == nil {
			Config.DbPort = port
		}
		Config.DbName = env("DB_NAME", Config.DbName)
		Config.DbUser = env("DB_USER", Config.DbUser)
		Config.DbPasswd = env("DB_Passwd", Config.DbPasswd)
		Config.DbCollation = env("DB_COLLATION", Config.DbCollation)
	}
}

func env(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}
