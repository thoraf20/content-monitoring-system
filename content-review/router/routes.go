package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/content-monitoring-system/content-review/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/review/pending", handler.ListPendingContent)
	r.POST("/review/:id", handler.ReviewContent)
}
