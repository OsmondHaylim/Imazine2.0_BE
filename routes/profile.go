package routes

import (
	"encoding/json"
	"fmt"
	"imazine/models"
	"imazine/storage"
	"imazine/utils"
	"io"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UploadProfilePicture(context *fiber.Ctx) error {
	file, err := context.FormFile("image")
	if err != nil {
		return context.Status(400).JSON(err.Error())
	}

	context.SaveFile(file, fmt.Sprintf("./download_cache/%s", file.Filename))

	values := map[string]io.Reader{
		"image": utils.MustOpen(fmt.Sprintf("download_cache/%s", file.Filename)), // lets assume its this file
	}

	res, err := utils.Upload("https://api.imgbb.com/1/upload?key=7b39ff8818a667ee516b470fd8bcbd09", values)

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return context.Status(400).JSON(err.Error())
	}
	bodyString := string(bodyBytes)

	type ExtractImgUrl struct {
		Data struct {
			Url string `json:"url"`
		}
	}

	var imgUrl ExtractImgUrl
	err = json.Unmarshal([]byte(bodyString), &imgUrl)

	err = os.Remove(fmt.Sprintf("download_cache/%s", file.Filename)) // remove a single file
	if err != nil {
		return context.Status(400).JSON(err.Error())
	}

	userLocals := context.Locals("user")
	user, _ := userLocals.(models.User)

	storage.DB.Db.Model(&user).Update("ProfilePictureLink", imgUrl.Data.Url)

	return context.SendStatus(200)
}

func ChangePassword(context *fiber.Ctx) error {
	type FormBody struct {
		OldPassword       string `form:"oldPassword"`
		NewPassword       string `form:"newPassword"`
		RepeatNewPassword string `form:"repeatNewPassword"`
	}

	a := new(FormBody)
	if err := context.BodyParser(a); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if a.NewPassword != a.RepeatNewPassword {
		return context.Status(400).JSON(&fiber.Map{
			"message": "Password baru tidak cocok!",
		})
	}

	userLocals := context.Locals("user")
	user, _ := userLocals.(models.User)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(a.OldPassword)); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": "NPM atau Password salah!",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.NewPassword), 10)
	if err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	storage.DB.Db.Model(&user).Update("Password", string(hashedPassword))

	return context.SendStatus(200)
}

func UpdateProfile(context *fiber.Ctx) error {
	userLocals := context.Locals("user")
	user, _ := userLocals.(models.User)

	type FormBody struct {
		Email string `form:"email"`
	}

	a := new(FormBody)
	if err := context.BodyParser(a); err != nil {
		return context.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	storage.DB.Db.Model(&user).Update("Email", a.Email)
	return context.SendStatus(200)
}
