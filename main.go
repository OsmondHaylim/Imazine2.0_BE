package main

import (
	"fmt"
	"imazine/models"
	"imazine/storage"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	// "encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
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
	DateCreated	string		`json:"created"`
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

	err := r.DB.Create(&article).Error
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

func(r *Repository) SetupRoutes(app *fiber.Ctx){
	api := app.Group("/api")
	api.Post("/create_articles", r.CreateArticle)
	api.Delete("/delete_articles", r.DeleteArticle)
	api.Get("/get_articles/:id", r.GetArticleByID)
	api.Get("/get_articles", r.GetArticle)
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

