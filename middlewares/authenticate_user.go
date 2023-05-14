package middlewares

import (
	"imazine/models"
	"imazine/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func AuthenticateUser(context *fiber.Ctx) error {
	res := strings.Split(context.GetReqHeaders()["Authorization"], " ")

	// Bentuk Header yang diinginkan:
	// Authorization: Bearer [token 20 digit]
	if len(res) != 2 && res[0] != "Bearer" && len(res[1]) != 20 {
		return context.Status(400).JSON(&fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := res[1]

	user := models.User{}
	if err := storage.DB.Db.Where("token = ?", token).Preload(clause.Associations).First(&user).Error; err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": "Unauthorized",
		})
	}

	context.Locals("user", user)

	return context.Next()
}