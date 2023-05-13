package models

type ArticleCategory struct {
	ID    int    `gorm:"primaryKey;autoIncrement;notNull" json:"id"`
	Name  string `gorm:"notNull" json:"name"`
	Users []User `gorm:"many2many:has_article_edit_access;" json:"users"`
}

// not for database table, but for response conversion to match API spec
type ArticleCategorySmall struct {
	ID   int    `gorm:"primaryKey;autoIncrement;notNull" json:"id"`
	Name string `gorm:"notNull" json:"name"`
}

func ToArticleCategorySmall(articleCategory ArticleCategory) ArticleCategorySmall {
	return ArticleCategorySmall{
		ID:   articleCategory.ID,
		Name: articleCategory.Name,
	}
}
