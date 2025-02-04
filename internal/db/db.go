package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"felipejsm/tp-admin/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
    isProd := os.Getenv("PROD")
    var dsn string
    if isProd == "" {
        config.LoadEnv()
        dsn = config.GetEnv("DATABASE_URL", "default")
    } else {
        dbHost := os.Getenv("DB_HOST")
        dbUser := os.Getenv("DB_USER")
        dbPassword := os.Getenv("DB_PASSWORD")
        dbName := os.Getenv("DB_NAME")
        dbPort := os.Getenv("DB_PORT")
        dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
        dbHost, dbUser, dbPassword, dbName, dbPort)
    }
   
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("[Error] database connection failed: %v", err)
	}
	sqlDB, err := DB.DB()

	if err != nil {
		log.Fatalf("[Error] sql db access failed: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("[InitDB] DB Connection Succesfully obtained.")
	return DB
}
