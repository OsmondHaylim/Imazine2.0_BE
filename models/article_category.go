package models

type ArticleCategory struct {
	ID    int    `gorm:"primaryKey;autoIncrement;notNull" json:"id"`
	Name  string `gorm:"notNull" json:"name"`
	Users []User `gorm:"many2many:has_article_edit_access;" json:"users"`
}