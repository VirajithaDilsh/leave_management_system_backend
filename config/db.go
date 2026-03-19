package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"localhost",        // DB_HOST
		"postgres",         // DB_USER
		"2328687",          // DB_PASSWORD
		"leave_management", // DB_NAME
		"5432",             // DB_PORT
		"disable",          // DB_SSLMODE
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")
}
