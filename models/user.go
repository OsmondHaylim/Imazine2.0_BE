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

// not for database table, but for response conversion to match API spec
type UserSmall struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	NPM                string `json:"npm"`
	ProfilePictureLink string `json:"profile_picture_link"`
}

func ToUserSmall(user User) UserSmall {
	return UserSmall{
		ID:                 user.ID,
		Name:               user.Name,
		NPM:                user.NPM,
		ProfilePictureLink: user.ProfilePictureLink,
	}
}
