package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "url-shortener/models"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open(sqlite.Open("urlshortener.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    DB.AutoMigrate(&models.URL{})
}