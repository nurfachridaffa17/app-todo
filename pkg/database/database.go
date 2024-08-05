package database

import (
	"app-todo/pkg/constant"
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections map[string]*gorm.DB
	database      *gorm.DB
)

func Init() {
	dbConfigurations := map[string]Db{
		constant.DB_NAME: &dbPostgreSQL{
			db: db{
				Host: os.Getenv(constant.DB_HOST),
				User: os.Getenv(constant.DB_USER),
				Pass: os.Getenv(constant.DB_PASS),
				Port: os.Getenv(constant.DB_PORT),
				Name: os.Getenv(constant.DB_NAME),
			},
			SslMode: os.Getenv(constant.DB_SSLMODE),
			Tz:      os.Getenv(constant.DB_TZ),
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		if k == constant.DB_NAME {
			database = db
		}
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func GetConnection(name string) (*gorm.DB, error) {
	if dbConnections[name] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[name], nil
}

func DBManager() *gorm.DB {
	return database
}
