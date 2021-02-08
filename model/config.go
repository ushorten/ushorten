package model

type Config struct {
	AppPath string
	AppDebug bool
	DbType string
	DbPath string
	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPasswd string
	DbCollation string
}

func NewConfig() *Config {
	config := &Config{
		AppDebug: false,
		DbType: "sqlite",
		DbPath: "data/ushorten.sqlite3",
		DbHost: "127.0.0.1",
		DbPort: 3306,
		DbName: "ushorten",
		DbUser: "root",
		DbCollation: "utf8mb4_unicode_ci",
	}
	config.DbPort = 3306
	return config
}
