package models

import (
	"time"
)

type Article struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Author          string          `gorm:"notNull" json:"author"`
	CategoryID      int             `gorm:"notNull" json:"category_id"`
	Category        ArticleCategory `json:"category"`
	Title           string          `gorm:"notNull" json:"title"`
	MarkdownContent string          `gorm:"notNull" json:"markdown_content "`
	CreatedAt       time.Time       `gorm:"notNull" json:"created_at"`
	ViewCount       int             `gorm:"notNull;default:0" json:"view_count"`
}
