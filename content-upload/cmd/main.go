package main

import (
	"github.com/thoraf20/content-monitoring-system/content-upload/internal/config"
	"github.com/thoraf20/content-monitoring-system/content-upload/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()
	r.Any("/upload", handler.HandleUpload)

	port := config.Get("PORT")
	if port == "" {
		port = "5002"
	}
	r.Run(":" + port)
}
