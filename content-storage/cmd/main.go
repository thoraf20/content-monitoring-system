package main

import (
	"github.com/thoraf20/content-monitoring-system/content-storage/config"
	"github.com/thoraf20/content-monitoring-system/content-storage/db"
	"github.com/thoraf20/content-monitoring-system/content-storage/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	r := gin.Default()
	router.SetupRoutes(r)

	port := config.Get("PORT")
	if port == "" {
		port = "5003"
	}
	r.Run(":" + port)
}
