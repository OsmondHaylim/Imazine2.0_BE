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

func MigrateArticles(db *gorm.DB) error{
	err := db.AutoMigrate(&Article{})
	return err
}