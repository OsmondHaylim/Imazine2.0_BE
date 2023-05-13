package models

type User struct {
	ID                   int               `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                 string            `gorm:"notNull" json:"name"`
	NPM                  string            `gorm:"notNull" json:"npm"`
	ProfilePictureLink   string            `json:"profile_picture_link"`
	Email                string            `gorm:"notNull" json:"email"`
	Password             string            `gorm:"notNull" json:"password"`
	IsAdmin              bool              `gorm:"notNull;default:false" json:"is_admin"`
	Token                string            `json:"token"`
	HasArticleEditAccess []ArticleCategory `gorm:"many2many:has_article_edit_access;" json:"has_article_edit_access"`
}