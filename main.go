package main

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
)

type Repository struct{
	DB *gorm.DB
}

// testing with yt tutor vid

type Article struct{
	ID 			string		`json:"id"`
	AuthorID 	string		`json:"author"`
	CategoryID 	string		`json:"category"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	DateCreated	string		`json:"created"`
	ViewCount	int			`json:"view_count"`
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

