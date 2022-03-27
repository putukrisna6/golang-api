package config

import (
	"fmt"
	"os"

	"github.com/putukrisna6/golang-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Receipt{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection")
	}
	dbSQL.Close()
}
