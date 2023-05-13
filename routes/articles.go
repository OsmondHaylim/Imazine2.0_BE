package routes

import (
	"fmt"
	"imazine/models"
	"imazine/storage"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)


func CreateArticle(context *fiber.Ctx) error{
	article := new(models.Article)

	if err := context.BodyParser(article); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Create(&article)

	var result models.Article

	storage.DB.Db.Preload(clause.Associations).First(&result, article.ID)

	return context.Status(200).JSON(&fiber.Map{
		"message": "Article created",
		"article": models.ToArticleSmall(result),
	})
}

func GetArticle(context *fiber.Ctx) error{
	artM := &[]models.Article{}

	err := storage.DB.Db.Find(artM).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"Could not get articles"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":	"Articles fetched successfully", 
		"data":		artM,
	})
	return err
}

func DeleteArticle(context *fiber.Ctx) error{
	artM := models.Article{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	err := storage.DB.Db.Delete(artM, id)

	if err.Error != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"Could not delete article",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Article deleted",
	})
	return nil
}

func GetArticleByID(context *fiber.Ctx) error{
	id := context.Params("id")
	artM := &models.Article{}
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	fmt.Println("ID = ", id)
	err := storage.DB.Db.Preload(clause.Associations).First(artM, id).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"Couldn't get article",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Article ID fetched",
		"data":artM,
	})
	return nil
}