package main

import (
	"log"
	"time"

	"leave-management-backend/config"
	"leave-management-backend/models"
	"leave-management-backend/routes"
	"leave-management-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	err := config.DB.AutoMigrate(&models.User{}, &models.LeaveRequest{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	var existingUser models.User
	config.DB.Where("email = ?", "admin@test.com").First(&existingUser)

	if existingUser.ID == 0 {
		hashedPassword, _ := utils.HashPassword("admin123")

		admin := models.User{
			Name:     "Admin User",
			Email:    "admin@test.com",
			Password: hashedPassword,
			Role:     "admin",
			JobTitle: "HR",
		}

		config.DB.Create(&admin)
		log.Println("Admin user seeded")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	r.Run(":8080")
}