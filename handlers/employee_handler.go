package handlers

import (
	"net/http"

	"leave-management-backend/config"
	"leave-management-backend/models"
	"leave-management-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

type CreateEmployeeInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	JobTitle string `json:"jobTitle"`
}

func CreateEmployee(c *gin.Context) {
	var input CreateEmployeeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not hash password",
		})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
		JobTitle: input.JobTitle,
	}

	if user.Role == "" {
		user.Role = "employee"
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not create employee",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee created successfully",
		"user": gin.H{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"role":     user.Role,
			"jobTitle": user.JobTitle,
		},
	})
}
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})
		return
	}

	// delete related leave requests first
	if err := config.DB.
		Where("employee_id = ?", user.ID).
		Delete(&models.LeaveRequest{}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete leave requests",
			"details": err.Error(),
		})
		return
	}

	// delete user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete employee",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee and leave records deleted",
	})
}