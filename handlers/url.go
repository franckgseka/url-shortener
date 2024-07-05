package handlers

import (
    "net/http"
    "url-shortener/database"
    "url-shortener/models"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type ShortenRequest struct {
    LongURL string `json:"long_url" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    shortURL := uuid.New().String()[:6]

    username, _ := c.Get("username")
    var user models.User
    if result := database.DB.Where("username = ?", username).First(&user); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
        return
    }

    url := models.URL{LongURL: req.LongURL, ShortURL: shortURL, UserID: user.ID}
    database.DB.Create(&url)

    c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func ResolveShortURL(c *gin.Context) {
    shortURL := c.Param("shortURL")
    var url models.URL
    if result := database.DB.First(&url, "short_url = ?", shortURL); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    // Incrémenter le compteur de clics
    database.DB.Model(&url).Update("clicks", url.Clicks+1)

    c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
