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
	config.LoadEnv()
	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))
	dsn := config.GetEnv("DATABASE_URL", "default")
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
