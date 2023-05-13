package main

import (
	"fmt"
	"imazine/models"
	"imazine/storage"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
	"log"
)

// testing with yt tutor vid

type Repository struct{
	DB *gorm.DB
}

type Article struct{
	Author 		string		`json:"author"`
	Category 	string		`json:"category"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	DateCreated	time.Time	`json:"created"`
	ViewCount	int			`json:"view_count"`
}

func (r *Repository) CreateArticle(context *fiber.Ctx) error{
	article := Article{}

	err := context.BodyParser(&article)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"Request failed"})
		return err
	}

	err = r.DB.Create(&article).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"Could not create article"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message":"Article has been added"})
	return nil
}
func (r *Repository) GetArticle(context *fiber.Ctx) error{
	artM := &[]models.Article{}

	err := r.DB.Find(artM).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"Could not get articles"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":	"Articles fetched successfully", 
		"data":		artM,
	})
	return err
}
func (r *Repository) DeleteArticle(context *fiber.Ctx) error{
	artM := models.Article{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(artM, id)

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
func (r *Repository) GetArticleByID(context *fiber.Ctx) error{
	id := context.Params("id")
	artM := &models.Article{}
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"ID cannot be empty",
		})
		return nil
	}

	fmt.Println("ID = ", id)
	err := r.DB.Where("id = ?", id).First(artM).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"Couldn't get article",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Article ID fetched",
		"data":artM,
	})
	return nil
}

func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_articles", r.CreateArticle)
	api.Delete("/delete_articles", r.DeleteArticle)
	api.Get("/get_articles/:id", r.GetArticleByID)
	api.Get("/get_articles", r.GetArticle)
	api.Post("/login", r.Login)
}

func(r *Repository) Login(context *fiber.Ctx) error{
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
			"user":    user,
		})
	}

	return c.Status(400).JSON(&fiber.Map{
		"message": "Login Failed!",
	})
}


func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:		os.Getenv("DB_HOST"),
		Port:		os.Getenv("DB_PORT"),
		Password:	os.Getenv("DB_PASS"),
		User:		os.Getenv("DB_USER"),
		SSLMode:	os.Getenv("DB_SSLMODE"),
		DBName:		os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = models.MigrateArticles(db)
	if err != nil {
		log.Fatal("Could not migrate the database")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
	// fmt.Println("starting web server at http://localhost:8080")
	// http.ListenAndServe(":8080", GetMux())
}

// func homepage() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//index.vue
// 	}
// }

// func article() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//article
// 	}
// }

// func manageArticle() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//Manage Article
// 	}
// }

// func profile() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//profile
// 	}
// }

// func login() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//login
// 	}
// }

// func register() http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//register
// 	}
// }



// func GetMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", homepage())
// 	return mux
// }

