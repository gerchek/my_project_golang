package gormclient

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// NewPgConfig creates new pg config instance
func NewGormConfig(username string, password string, host string, port string, database string) *gormConfig {
	return &gormConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

// NewClient
func NewClient(cfg *gormConfig) (client *gorm.DB, err error) {
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Database)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.Book{})

	return db, err
}

//CloseDatabaseConnection method is closing a connection between app and database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
