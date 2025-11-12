package routes

import (
	"laporan-lingkungan/controller"
	"laporan-lingkungan/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	reports := api.Group("/reports")
	reports.Use(middleware.AuthMiddleware())
	{
		reports.POST("", controllers.CreateReport)
		reports.GET("", controllers.GetAllReports)
		reports.GET("/my-reports", controllers.GetMyReports)
		reports.GET("/stats", controllers.GetReportStats)
		reports.PUT("/:id", controllers.UpdateReport)
		reports.DELETE("/:id", controllers.DeleteReport)
	}
}
