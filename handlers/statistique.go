package handlers

import (
	"net/http"
	"url-shortener/database"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

func GetStatistics(c *gin.Context) {
	var urls []models.URL
	database.DB.Find(&urls)

	totalLinks := len(urls)
	totalClicks := 0
	for _, url := range urls {
		totalClicks += url.Clicks
	}

	c.JSON(http.StatusOK, gin.H{
		"total_links":  totalLinks,
		"total_clicks": totalClicks,
		"details":      urls,
	})
}
