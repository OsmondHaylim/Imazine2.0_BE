package routes

import (
	"encoding/json"
	"fmt"
	"imazine/models"
	"imazine/storage"
	"imazine/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"

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
	return nil
}

func GetRequestByID(context *fiber.Ctx) error{
	return nil
}