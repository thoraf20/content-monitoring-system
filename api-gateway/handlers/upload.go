package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		// Save file locally (just for dev testing)
		err := c.SaveUploadedFile(file, fmt.Sprintf("uploads/%s", file.Filename))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "upload received"})
}
