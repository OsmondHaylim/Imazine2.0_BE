package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"imazine/models"
	"imazine/storage"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// TODO clean up the whole upload thing, maybe move this to another file and move stuffs on CreateArticle to separate functions
func Upload(url string, values map[string]io.Reader) (res *http.Response, err error) {
	var client *http.Client = &http.Client{}
	
    // Prepare a form that you will submit to that URL.
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    for key, r := range values {
        var fw io.Writer
        if x, ok := r.(io.Closer); ok {
            defer x.Close()
        }
        // Add an image file
        if x, ok := r.(*os.File); ok {
            if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
                return
            }
        } else {
            // Add other fields
            if fw, err = w.CreateFormField(key); err != nil {
                return
            }
        }
        if _, err = io.Copy(fw, r); err != nil {
            return
        }

    }
    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    w.Close()

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return nil, err
    }

    // Don't forget to set the content type, this will contain the boundary.
    req.Header.Set("Content-Type", w.FormDataContentType())
    req.Header.Set("Content-Length", fmt.Sprint(len(b.Bytes())))
	req.Header.Set("Host", "localhost:8080")

    // Submit the request
    res, err = client.Do(req)
	return res, err
}

func mustOpen(f string) *os.File {
    r, err := os.Open(f)
    if err != nil {
        panic(err)
    }
    return r
}


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
        "image": mustOpen(fmt.Sprintf("download_cache/%s", file.Filename)), // lets assume its this file
    }
    
	res, err := Upload("https://api.imgbb.com/1/upload?key=7b39ff8818a667ee516b470fd8bcbd09", values)
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