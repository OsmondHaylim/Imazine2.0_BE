package main

import (
	"context"

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

func (r *Repository) CreateArticle(context *fiber.App) error{
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

func (r *Repository) GetArticle(context *fiber.App) error{
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



func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_articles", r.CreateArticle)
	api.Delete("/delete_articles", r.DeleteArticle)
	api.Get("/get_articles/:id", r.GetArticleByID)
	api.Get("/get_articles", r.GetArticle)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database")
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

