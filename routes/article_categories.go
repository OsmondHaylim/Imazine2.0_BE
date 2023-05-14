package routes

import (
	"imazine/models"
	"imazine/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateArticleCategory(context *fiber.Ctx) error{
	articleCategory := new(models.ArticleCategory)

	if err := context.BodyParser(articleCategory); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Create(&articleCategory)

	return context.Status(200).JSON(models.ToArticleCategorySmall(*articleCategory))
}

func GetArticleCategories(context *fiber.Ctx) error {
	var articleCategories []models.ArticleCategorySmall

	result := storage.DB.Db.Model(&models.ArticleCategory{}).Find(&articleCategories)
	if result.Error != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return context.Status(200).JSON(articleCategories)
}

func GetArticleCategoryByID(context *fiber.Ctx) error {
	id := context.Params("id")
	var articleCategory models.ArticleCategorySmall

	result := storage.DB.Db.Model(&models.ArticleCategory{}).First(&articleCategory, id)
	if result.Error != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return context.Status(200).JSON(articleCategory)
}

func EditCategory(context *fiber.Ctx) error {
	articleCategory := new(models.ArticleCategory)
	i, _ := strconv.Atoi(context.Params("id"))

	if err := context.BodyParser(articleCategory); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	articleCategory.ID = i
	storage.DB.Db.Save(&articleCategory)

	return context.Status(200).JSON(models.ToArticleCategorySmall(*articleCategory))
}

// TODO handle article FK constraint
func DeleteCategory(context *fiber.Ctx) error {
	id := context.Params("id")
	var articleCategory models.ArticleCategorySmall

	result := storage.DB.Db.Model(&models.ArticleCategory{}).First(&articleCategory, id)
	if result.Error != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": result.Error.Error(),
		})
	}

	result = storage.DB.Db.Delete(&models.ArticleCategory{}, id)

	if result.Error != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return context.Status(200).JSON(articleCategory)
}