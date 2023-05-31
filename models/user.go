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
	IsAdmin            bool   `json:"is_admin"`
}

type UserLogin struct {
	ID                   int               `json:"id"`
	Name                 string            `json:"name"`
	NPM                  string            `json:"npm"`
	ProfilePictureLink   string            `json:"profile_picture_link"`
	Email                string            `json:"email"`
	IsAdmin              bool              `json:"is_admin"`
	HasArticleEditAccess []ArticleCategory `json:"has_article_edit_access"`
}

func ToUserSmall(user User) UserSmall {
	return UserSmall{
		ID:                 user.ID,
		Name:               user.Name,
		NPM:                user.NPM,
		ProfilePictureLink: user.ProfilePictureLink,
	}
}

func ToUserLogin(user User) UserLogin {
	return UserLogin{
		ID:                   user.ID,
		Name:                 user.Name,
		NPM:                  user.NPM,
		ProfilePictureLink:   user.ProfilePictureLink,
		Email:                user.Email,
		IsAdmin:              user.IsAdmin,
		HasArticleEditAccess: user.HasArticleEditAccess,
	}
}

func CreateUserList(data [][]string) []User {
	var users []User
	for i, line := range data {
		if i > 0 { // omit header line
			var rec User
			for j, field := range line {
				if j == 0 {
					rec.NPM = field
				} else if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Email = field
				}
			}
			users = append(users, rec)
		}
	}
	return users
}
