package models

import (
    "gorm.io/gorm"
)

type URL struct {
    gorm.Model
    LongURL  string `json:"long_url"`
    ShortURL string `json:"short_url"`
    Clicks   int    `json:"clicks"`
    UserID   uint   `json:"user_id"`
    User     User   `json:"user" gorm:"foreignKey:UserID"`
}
