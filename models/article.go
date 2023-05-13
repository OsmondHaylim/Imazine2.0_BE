package models

import "gorm.io/gorm"

type Article struct{
	ID 			int			`gorm:"primary key;autoIncrement" json:"`
	Author 		*string		`json:"author"`
	Category 	*string		`json:"category"`
	Title		*string		`json:"title"`
	Content		*string		`json:"content"`
	DateCreated	*string		`json:"created"`
	ViewCount	*int		`json:"view_count"`
}

type FormBody struct {
	NPM      string `form:"npm"`
	Password string `form:"password"`
}

type User struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	NPM                  string `json:"npm"`
	ProfilePictureLink   string `json:"profile_picture_link"`
	Email                string `json:"email"`
	IsAdmin              bool   `json:"is_admin"`
	HasArticleEditAccess []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"has_article_edit_access"`
}

func MigrateArticles(db *gorm.DB) error{
	err := db.AutoMigrate(&Article{})
	return err
}