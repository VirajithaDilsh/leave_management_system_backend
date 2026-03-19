package routes

import (
	"leave-management-backend/handlers"
	"leave-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/auth/login", handlers.Login)

		api.GET("/employees", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.GetEmployees)
		api.POST("/employees", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.CreateEmployee)
		api.DELETE("/employees/:id", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.DeleteEmployee)
		api.PUT("/auth/change-password", middleware.AuthMiddleware(), handlers.ChangePassword)

		api.POST("/leaves", middleware.AuthMiddleware(), middleware.AuthorizeRole("employee", "admin"), handlers.CreateLeaveRequest)
		api.GET("/leaves/my", middleware.AuthMiddleware(), handlers.GetMyLeaves)
		api.GET("/leaves", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.GetAllLeaves)
		api.PUT("/leaves/:id/approve", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.ApproveLeave)
		api.PUT("/leaves/:id/reject", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), handlers.RejectLeave)
	}
}