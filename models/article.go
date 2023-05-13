package models

import (
	"time"
)

type Article struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatorID       int             `gorm:"notNull" json:"creator_id" form:"creator_id"`
	Creator         User            `gorm:"notNull;foreignKey:CreatorID" json:"creator"`
	CategoryID      int             `gorm:"notNull" json:"category_id" form:"category_id"`
	Category        ArticleCategory `gorm:"notNull;foreignKey:CategoryID" json:"category"`
	Title           string          `gorm:"notNull" json:"title" form:"title"`
	CoverImageLink  string          `gorm:"notNull" json:"cover_image_link"`
	MarkdownContent string          `gorm:"notNull" json:"markdown_content" form:"markdown_content"`
	CreatedAt       time.Time       `gorm:"notNull" json:"created_at"`
	ViewCount       int             `gorm:"notNull;default:0" json:"view_count"`
}

// not for database table, but for response conversion to match API spec
type ArticleSmall struct {
	ID              int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Creator         UserSmall       `gorm:"notNull;foreignKey:CreatorID" json:"creator"`
	Category        ArticleCategorySmall `gorm:"notNull;foreignKey:CategoryID" json:"category"`
	Title           string          `gorm:"notNull" json:"title" form:"title"`
	CoverImageLink  string          `gorm:"notNull" json:"cover_image_link"`
	MarkdownContent string          `gorm:"notNull" json:"markdown_content" form:"markdown_content"`
	CreatedAt       time.Time       `gorm:"notNull" json:"created_at"`
	ViewCount       int             `gorm:"notNull;default:0" json:"view_count"`
}

func ToArticleSmall(article Article) ArticleSmall {
	return ArticleSmall{
		ID:              article.ID,
		Creator:         ToUserSmall(article.Creator),
		Category:        ToArticleCategorySmall(article.Category),
		Title:           article.Title,
		CoverImageLink:  article.CoverImageLink,
		MarkdownContent: article.MarkdownContent,
		CreatedAt:       article.CreatedAt,
		ViewCount:       article.ViewCount,
	}
}
