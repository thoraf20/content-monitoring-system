package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/content-monitoring-system/content-review/router"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(":5004") // Runs on port 5004
}
