package routes

import (
	"imazine/models"
	"imazine/storage"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetArticleByCategory(context *fiber.Ctx) error{
	articles := &[]models.Article{}

	category := context.Params("id")

	err := storage.DB.Db.Preload(clause.Associations).Order("created_at desc").Where("category_id = ?", category).Find(&articles).Error
	if err != nil {
		return context.Status(400).JSON(err)
	}

	return context.Status(200).JSON(articles)
}