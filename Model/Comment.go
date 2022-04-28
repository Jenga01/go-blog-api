package models

import (
	"time"
)

type Comment struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"ID"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:10000;not null;" json:"content"`
	ArticleID uint64    `gorm:"not null foreignKey:ArticleID" json:"article_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
