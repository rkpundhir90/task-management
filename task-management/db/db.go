package db

import (
	"fmt"
	"log"

	"github.com/rkpundhir90/task-management/task-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database using GORM
func InitDB(cfg config.EnvironmentConfiguration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected!")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
