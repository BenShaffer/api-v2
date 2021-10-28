package infrastructure

import (
	"api/internal/domain"
	"api/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	ConnectionString string
}

type SQLDatabase struct {
	logger log.IApiLogger
	*gorm.DB
}

func NewSQLDatabase(logger log.IApiLogger, config Config) *SQLDatabase {
	db, err := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})

	if err != nil {
		logger.Panicf("Failed to connect to DB with error: %s", err.Error())
		return nil
	}

	return &SQLDatabase{
		DB:     db,
		logger: logger,
	}
}

func (db *SQLDatabase) Migrate() {
	db.logger.Debug("Migrating database...")

	if err := db.AutoMigrate(
		&domain.Person{},
	); err != nil {
		db.logger.Panicf("Failed to migrate DB with error: %s", err.Error())
		return
	}

	db.logger.Debug("Database migrated.")
}
