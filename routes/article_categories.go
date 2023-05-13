package routes

import (
	"imazine/models"
	"imazine/storage"

	"github.com/gofiber/fiber/v2"
)

func CreateArticleCategory(context *fiber.Ctx) error{
	articleCategory := new(models.ArticleCategory)

	if err := context.BodyParser(articleCategory); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Create(&articleCategory)

	return context.Status(200).JSON(&fiber.Map{
		"message": "Article category created",
		"article_category": articleCategory,
	})
}