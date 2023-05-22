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

	
}

func UpdateRequest(context *fiber.Ctx) error {
	return nil
}

func GetRequest(context *fiber.Ctx) error{
	return nil
}

func DeleteRequest(context *fiber.Ctx) error{
	return nil
}

func GetRequestByID(context *fiber.Ctx) error{
	return nil
}