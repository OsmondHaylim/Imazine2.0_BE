package models

import "time"

type Request struct{
	ID 				int 			`gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID      int             `gorm:"notNull" json:"category_id" form:"category_id"`
	Category        ArticleCategory `gorm:"notNull;foreignKey:CategoryID" json:"category"`
	Content 		string          `gorm:"notNull" json:"content" form:"content"`
	Status 			string			`gorm:"notNull" json:"status"`
	CreatedAt   	time.Time       `gorm:"notNull" json:"created_at"`
	Deadline		time.Time		`gorm:"notNull" json:"deadline"`
}