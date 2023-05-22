package models

import "time"

type Request struct{
	ID 				int 			`gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID      int             `gorm:"notNull" json:"category_id" form:"category_id"`
	Category        ArticleCategory `gorm:"notNull;foreignKey:CategoryID" json:"category"`
	ContentHeader 	string          `gorm:"notNull" json:"content_header" form:"content_header"`
	Content 		string          `gorm:"notNull" json:"content" form:"content"`
	Status 			string			`gorm:"notNull" json:"status"`
	CreatedAt   	time.Time       `gorm:"notNull" json:"created_at"`
	Deadline		time.Time		`gorm:"notNull" json:"deadline"`
}
type RequestSmall struct{
	ID 				int 			`gorm:"primaryKey;autoIncrement" json:"id"`
	ContentHeader 	string          `gorm:"notNull" json:"content_header" form:"content_header"`
	Status 			string			`gorm:"notNull" json:"status"`
	Deadline		time.Time		`gorm:"notNull" json:"deadline"`
}

func ToRequestSmall(request Request) RequestSmall {
	return RequestSmall{
		ID:              	request.ID,
		ContentHeader:   	request.ContentHeader,
		Status:        		request.Status,
		Deadline:           request.Deadline,
	}
}