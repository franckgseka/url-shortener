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

    r.POST("/signup", handlers.SignUp)
    r.POST("/login", handlers.Login)

    authorized := r.Group("/")
    authorized.Use(handlers.AuthMiddleware())
    {
        authorized.POST("/shorten", handlers.CreateShortURL)
        authorized.GET("/stats", handlers.GetStatistics)
    }

    r.GET("/:shortURL", handlers.ResolveShortURL)

    r.Run(":8080")
}
