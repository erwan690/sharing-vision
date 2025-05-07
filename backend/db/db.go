package db

import (
	"fmt"
	"strings"
	"sync"

	"github.com/erwan690/blog/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	*gorm.DB
}

var (
	dbConn *DatabaseConnection
	once   sync.Once
)

func GetDatabaseConnection(config *config.Config) (*DatabaseConnection, error) {
	var initError error
	once.Do(func() {
		dsn := strings.TrimPrefix(config.DBURL, "mysql://")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			initError = fmt.Errorf("failed to open database: %v", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			initError = fmt.Errorf("failed to use database: %v", err)
			return
		}

		if err = sqlDB.Ping(); err != nil {
			initError = fmt.Errorf("failed to ping database: %v", err)
			return
		}
		dbConn = &DatabaseConnection{db}
	})
	if initError != nil {
		return nil, initError
	}
	return dbConn, nil
}
