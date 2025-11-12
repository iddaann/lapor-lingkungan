package main

import (
	"laporan-lingkungan/config"
	"laporan-lingkungan/middleware"
	"laporan-lingkungan/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	config.ConnectDatabase()
	middleware.AutoMigrate()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
