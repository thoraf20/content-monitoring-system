package handler

import (
	"github.com/thoraf20/content-monitoring-system/content-storage/db"
	"github.com/thoraf20/content-monitoring-system/content-storage/model"
	"github.com/thoraf20/content-monitoring-system/content-storage/storage"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"os"
)

func HandleStore(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	storagePath := os.Getenv("STORAGE_PATH")
	savedPath, err := storage.SaveFile(file, storagePath)
	if err != nil {
		log.Println("File saving error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	content := model.Content{
		FileName:   file.Filename,
		FileType:   filepath.Ext(file.Filename),
		Path:       savedPath,
		UploadedAt: time.Now(),
	}

	if err := db.DB.Create(&content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content stored successfully", "content": content})
}

func ListFiles(c *gin.Context) {
	var contents []model.Content
	if err := db.DB.Find(&contents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content"})
		return
	}
	c.JSON(http.StatusOK, contents)
}

func GetFile(c *gin.Context) {
	id := c.Param("id")

	var content model.Content
	if err := db.DB.First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	filePath := filepath.Join(os.Getenv("STORAGE_PATH"), content.FileName)
	c.File(filePath)
}
