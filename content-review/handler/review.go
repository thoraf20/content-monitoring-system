package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/content-monitoring-system/content-review/model"
)

var reviewQueue = []model.Content{
	{ID: "1", Type: "image", URL: "http://cdn.com/image1.jpg", Status: "pending"},
	{ID: "2", Type: "text", URL: "http://cdn.com/text1.txt", Status: "pending"},
}

func ListPendingContent(c *gin.Context) {
	var pending []model.Content
	for _, item := range reviewQueue {
		if item.Status == "pending" {
			pending = append(pending, item)
		}
	}
	c.JSON(http.StatusOK, pending)
}

func ReviewContent(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status  string `json:"status"` // "approved" or "rejected"
		Comment string `json:"comment"`
		User    string `json:"reviewed_by"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range reviewQueue {
		if item.ID == id {
			reviewQueue[i].Status = input.Status
			reviewQueue[i].ReviewedBy = input.User
			reviewQueue[i].Comment = input.Comment
			c.JSON(http.StatusOK, reviewQueue[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "content not found"})
}
