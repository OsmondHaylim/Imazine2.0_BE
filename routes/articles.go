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

func CreateArticle(context *fiber.Ctx) error{
	article := new(models.Article)

	if err := context.BodyParser(article); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	file, err := context.FormFile("cover_image")
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


	article.CreatorID = user.ID
	article.CoverImageLink = imgUrl.Data.Url

	err = storage.DB.Db.Create(&article).Error
	if err != nil {
		return context.Status(400).JSON(err.Error())
	}

	var result models.Article
	storage.DB.Db.Preload(clause.Associations).First(&result, article.ID)

	return context.Status(200).JSON(&fiber.Map{
		"message": "Article created",
		"article": models.ToArticleSmall(result),
	})
}

func GetArticle(context *fiber.Ctx) error{
	articles := &[]models.Article{}
	query := storage.DB.Db.Preload(clause.Associations).Order("created_at desc")

	if categoryId := context.Query("category"); categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}

	// TODO implement limit and offset
	// TODO only fetch needed data (articleListCard)
	err := query.Find(articles).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(err)
	}

	return context.Status(http.StatusOK).JSON(articles)
}

func DeleteArticle(context *fiber.Ctx) error{
	artM := models.Article{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	err := storage.DB.Db.Delete(artM, id)

	if err.Error != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"Could not delete article",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Article deleted",
	})
	return nil
}

func GetArticleByID(context *fiber.Ctx) error{
	id := context.Params("id")
	artM := &models.Article{}

	err := storage.DB.Db.Preload(clause.Associations).First(artM, id).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(err)
	}
	context.Status(http.StatusOK).JSON(models.ToArticleSmall(*artM))
	return nil
}