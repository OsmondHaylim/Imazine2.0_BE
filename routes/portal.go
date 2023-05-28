package routes

import (
	// "encoding/json"
	// "fmt"
	"imazine/models"
	"imazine/storage"
	"imazine/utils"
	// "io"
	// "io/ioutil"
	// "net/http"
	// "os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func CreateRequest(context *fiber.Ctx) error{
	request := new(models.Request)

	if err := context.BodyParser(request); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	return context.Status(200).JSON(models.ToRequestSmall(request))
}

func UpdateRequest(context *fiber.Ctx) error {
	request := new(models.Request)

	storage.DB.Db.First(request, context.Params("id"))

	if err := context.BodyParser(request); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	storage.DB.Db.Save(&request)

	return context.Status(200).JSON(models.ToRequestSmall(request))
}

func GetRequest(context *fiber.Ctx) error{
	requests := &[]models.Request{}
	query := storage.DB.Db.Preload(clause.Associations).Order("deadline asc")

	err := query.Find(requests).Error
	if err != nil {
		return context.Status(400).JSON(err)
	}

	return context.Status(200).JSON(requests)
}

func DeleteRequest(context *fiber.Ctx) error{
	reqM := models.Request{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	err := storage.DB.Db.Delete(reqM, id)

	if err.Error != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"Could not delete request",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Request deleted",
	})
	return nil
}

func GetRequestByID(context *fiber.Ctx) error{
	id := context.Params("id")
	reqM := &models.Request{}

	err := storage.DB.Db.Preload(clause.Associations).First(reqM, id).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(err)
	}
	context.Status(http.StatusOK).JSON(models.ToRequestSmall(*reqM))
	return nil
}