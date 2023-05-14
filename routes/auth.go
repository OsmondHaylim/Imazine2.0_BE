package routes

import (
	"crypto/rand"
	"encoding/hex"
	"imazine/models"
	"imazine/storage"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func Login(c *fiber.Ctx) error{
	type FormBody struct {
		NPM      string `form:"npm"`
		Password string `form:"password"`
	}
	a := new(FormBody)
	if err := c.BodyParser(a); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User
	if err := storage.DB.Db.Where("npm = ?", a.NPM).Preload(clause.Associations).First(&user).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "NPM atau Password salah!",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(a.Password)); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "NPM atau Password salah!",
		})
	}

	token := GenerateSecureToken(10)
	storage.DB.Db.Model(&user).Update("Token", token)

	return c.Status(200).JSON(&fiber.Map{
		"message": "Login Success!",
		"apiKey": token,
		"user":    models.ToUserLogin(user),
	})
}

// probably not gonna be exposed to frontend
func Register(context *fiber.Ctx) error {
	userModel := new(models.User)

	if err := context.BodyParser(userModel); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := DoRegister(userModel); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	
	return context.Status(200).JSON(userModel)
}

func DoRegister(userModel *models.User)  error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 10)

	userModel.Password = string(newPassword)

	storage.DB.Db.Create(&userModel)

	return err
}