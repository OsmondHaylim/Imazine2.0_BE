package main

import (
	"encoding/json"
	"imazine/routes"
	"imazine/storage"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
)

// testing with yt tutor vid


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

func setUpRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Post("/categories", routes.CreateArticleCategory)

	app.Get("/articles", routes.GetArticle)
	app.Get("/articles/:id", routes.GetArticleByID)
	app.Post("/articles", routes.CreateArticle)
	app.Delete("/articles", routes.DeleteArticle)
	app.Post("/login", Login)
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

	storage.ConnectDB(config)
	app := fiber.New()
	
	setUpRoutes(app)


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

