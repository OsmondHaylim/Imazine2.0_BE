package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
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