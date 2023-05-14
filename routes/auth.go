package routes

import (
	"encoding/json"
	"imazine/models"
	"imazine/storage"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error{
	type FormBody struct {
		NPM      string `form:"npm"`
		Password string `form:"password"`
	}
	a := new(FormBody)

	err := c.BodyParser(a)
	if err != nil {
		panic(err)
	}

	type ExportedUser struct {
		ID                   string `json:"id"`
		Name                 string `json:"name"`
		NPM                  string `json:"npm"`
		ProfilePictureLink   string `json:"profile_picture_link"`
		Email                string `json:"email"`
		IsAdmin              bool   `json:"is_admin"`
		HasArticleEditAccess []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"has_article_edit_access"`
	}

	jsonStr := `{
		"id": "aaaaaa",
		"name": "Fauzan Azmi Dwicahyo",
		"npm": "140810200030",
		"profile_picture_link": "https://a.ppy.sh/2449200?1624766977.jpeg",
		"email": "fauzan.azmi01@gmail.com",
		"is_admin": true,
		"has_article_edit_access": [
			{
				"id": 1,
				"name": "Big Category"
			}
		]
	}`

	var user ExportedUser
	err = json.Unmarshal([]byte(jsonStr), &user)

	if a.NPM == "npm" && a.Password == "password" {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Login Success!",
			"apiKey": "1234567890",
			"user":    user,
		})
	}

	return c.Status(400).JSON(&fiber.Map{
		"message": "Login Failed!",
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