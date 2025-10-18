package database

import (
	"fmt"
	"log"

	"github.com/alwijein/be_football/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase establishes connection to MySQL database
func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	// MySQL DSN format: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database MySQL connected successfully")
	DB = db
	return db, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
