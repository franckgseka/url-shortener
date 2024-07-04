package main

import (
    "url-shortener/database"
    "url-shortener/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize the database
    database.Init()

    r.POST("/shorten", handlers.CreateShortURL)
    r.GET("/:shortURL", handlers.ResolveShortURL)
    r.GET("/stats", handlers.GetStatistics)

    r.Run(":8080")
}
