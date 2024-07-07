package infrastructure

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	User      string
	Password  string
	Host      string
	Port      string
	Net       string
	Database  string
	ParseTime bool
}

func GetConfig() Config {
	return Config{
		MySQL: MySQLConfig{
			User:      os.Getenv("MYSQL_ROOT_USER"),
			Password:  os.Getenv("MYSQL_ROOT_PASSWORD"),
			Host:      os.Getenv("MYSQL_HOST"),
			Port:      os.Getenv("MYSQL_PORT"),
			Net:       "tcp", // SSH 接続を行う場合「mysql+tcp」にする必要あり
			Database:  os.Getenv("MYSQL_DATABASE"),
			ParseTime: true,
		},
	}
}

func (c MySQLConfig) DataSource() string {
	config := &mysql.Config{
		User:      c.User,
		Passwd:    c.Password,
		Addr:      fmt.Sprintf("%s:%s", c.Host, c.Port),
		Net:       c.Net,
		DBName:    c.Database,
		ParseTime: c.ParseTime,
	}
	return config.FormatDSN()
}
