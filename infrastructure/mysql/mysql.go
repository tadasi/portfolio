package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/tadasi/portfolio/infrastructure"
)

var driver = "mysql"

func Open() (*sqlx.DB, error) {
	config := infrastructure.GetConfig()
	db, err := sqlx.Open(driver, config.MySQL.DataSource())
	if err != nil {
		return nil, err
	}
	return db, nil
}
