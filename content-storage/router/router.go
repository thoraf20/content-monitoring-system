package router

import (
	"github.com/thoraf20/content-monitoring-system/content-storage/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/store", handler.HandleStore)
	r.GET("/files", handler.ListFiles)
	r.GET("/file/:id", handler.GetFile)
}
