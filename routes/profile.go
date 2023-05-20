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

	err = os.Remove(fmt.Sprintf("download_cache/%s", file.Filename))  // remove a single file
	if err != nil {
		return context.Status(400).JSON(err.Error())
	}

	userLocals := context.Locals("user") 
	user, _ := userLocals.(models.User)

	storage.DB.Db.Model(&user).Update("ProfilePictureLink", imgUrl.Data.Url)

	return context.SendStatus(200)
}