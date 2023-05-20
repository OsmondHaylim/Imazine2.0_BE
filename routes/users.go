package routes

import (
	"imazine/models"
	"imazine/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SearchUser(context *fiber.Ctx) error {
	query := storage.DB.Db.Model(&models.User{})

	if search := context.Query("search"); len(search) >= 3 {
		query = query.Where("lower(name) LIKE ?", "%" + strings.ToLower(search) + "%").
					  Or("npm LIKE ?", "%" + search + "%")
	}

	users := &[]models.UserSmall{}

	if err := query.Order("npm asc").Find(users).Error; err != nil {
		return context.Status(400).JSON(err)
	}

	return context.Status(200).JSON(users)
}