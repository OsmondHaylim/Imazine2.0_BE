package routes

import (
	"imazine/models"
	"imazine/storage"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetUsersWithCategoryEditAccess(context *fiber.Ctx) error {
	categoryId := context.Query("category")

	category := &models.ArticleCategory{}
	storage.DB.Db.Preload(clause.Associations).Find(category, categoryId)

	var users []models.UserSmall

	for _, elem := range category.Users {
		users = append(users, models.ToUserSmall(elem))
	}

	return context.Status(200).JSON(users)
}

func AddCategoryUserPair(context *fiber.Ctx) error {
	type Body struct {
		ArticleCategoryId int `form:"category_id"`
		UserId int `form:"user_id"`
	}

	body := new(Body)
	if err := context.BodyParser(body); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Table("has_article_edit_access").Create(body)
	return context.Status(200).JSON(body)
}

func RemoveCategoryUserPair(context *fiber.Ctx) error {
	type Body struct {
		ArticleCategoryId int `form:"category_id"`
		UserId int `form:"user_id"`
	}

	body := new(Body)
	if err := context.BodyParser(body); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.
		Table("has_article_edit_access").
		Where("user_id = ?", body.UserId).
		Where("article_category_id = ?", body.ArticleCategoryId).
		Delete(body)
	return context.Status(200).JSON(body)
}

func AddUser(context *fiber.Ctx) error {
	userModel := new(models.User)

	if err := context.BodyParser(userModel); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	userModel.Password = "defaultPassword"

	if err := DoRegister(userModel); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	
	return context.SendStatus(200)
}

func EditUser(context *fiber.Ctx) error {
	user := new(models.User)

	storage.DB.Db.First(user, context.Params("id"))

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Save(&user)
	return context.SendStatus(200)
}