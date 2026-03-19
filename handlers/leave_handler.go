package handlers

import (
	"net/http"

	"leave-management-backend/config"
	"leave-management-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateLeaveRequest(c *gin.Context) {
	var leave models.LeaveRequest

	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	leave.EmployeeID = userID.(uint)
	leave.Status = "pending"

	if err := config.DB.Create(&leave).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create leave request"})
		return
	}

	c.JSON(http.StatusCreated, leave)
}

func GetMyLeaves(c *gin.Context) {
	userID, _ := c.Get("userID")

	var leaves []models.LeaveRequest
	config.DB.Where("employee_id = ?", userID).Find(&leaves)

	c.JSON(http.StatusOK, leaves)
}

func GetAllLeaves(c *gin.Context) {
	var leaves []models.LeaveRequest
	config.DB.Preload("Employee").Find(&leaves)
	c.JSON(http.StatusOK, leaves)
}

func ApproveLeave(c *gin.Context) {
	id := c.Param("id")
	var leave models.LeaveRequest

	if err := config.DB.First(&leave, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave request not found"})
		return
	}

	leave.Status = "approved"
	config.DB.Save(&leave)

	c.JSON(http.StatusOK, gin.H{"message": "Leave approved"})
}

func RejectLeave(c *gin.Context) {
	id := c.Param("id")
	var leave models.LeaveRequest

	if err := config.DB.First(&leave, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave request not found"})
		return
	}

	leave.Status = "rejected"
	config.DB.Save(&leave)

	c.JSON(http.StatusOK, gin.H{"message": "Leave rejected"})
}