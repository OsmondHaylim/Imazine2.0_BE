package routes

import (
	"imazine/models"
	"imazine/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SearchUser(context *fiber.Ctx) error {
	search := context.Query("search")
	if len(search) < 3 {
		return context.Status(200).JSON([]string{})
	}
	
	users := &[]models.UserSmall{}

	query := storage.DB.Db.Model(&models.User{})

	query = query.Where("lower(name) LIKE ?", "%" + strings.ToLower(search) + "%").
				  Or("npm LIKE ?", "%" + search + "%")

	if err := query.Find(users).Error; err != nil {
		return context.Status(400).JSON(err)
	}

	return context.Status(200).JSON(users)
}