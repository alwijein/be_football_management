package database

import (
	"log"

	"github.com/alwijein/be_football/internal/domain"
	"gorm.io/gorm"
)

// AutoMigrate runs auto migration for all models
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&domain.User{},
		&domain.Team{},
		&domain.Player{},
		&domain.Match{},
		&domain.Goal{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
