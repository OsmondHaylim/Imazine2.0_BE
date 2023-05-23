package main

import (
	"imazine/middlewares"
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

func setUpRoutes(app *fiber.App) {
	app.Use(cors.New())

	// route di atas app.Use(AuthenticateUser) nggak kena cek auth
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)

	app.Use(middlewares.AuthenticateUser)
	// Untuk mengakses user yang melakukan request, bisa menggunakan context.Locals("user")

	app.Post("/categories", routes.CreateArticleCategory)
	app.Get("/categories", routes.GetArticleCategories)
	app.Get("/categories/:id", routes.GetArticleCategoryByID)
	app.Put("/categories/:id", routes.EditCategory)
	app.Delete("/categories/:id", routes.DeleteCategory)

	app.Get("/articles", routes.GetArticle)
	app.Get("/articles/:id", routes.GetArticleByID)
	app.Post("/articles", routes.CreateArticle)
	app.Put("/articles/:id", routes.UpdateArticle)
	app.Delete("/articles/:id", routes.DeleteArticle)

	app.Get("/creators/articles/:id", routes.GetArticleByCategory)

	app.Get("/me", routes.Me)
	app.Get("/users", routes.SearchUser)
	app.Get("admin/categories/edit-access", routes.GetUsersWithCategoryEditAccess)
	app.Post("admin/categories/edit-access", routes.AddCategoryUserPair)
	app.Delete("admin/categories/edit-access", routes.RemoveCategoryUserPair)

	app.Put("/profile", routes.UpdateProfile)
	app.Post("/profile/password", routes.ChangePassword)
	app.Post("/profile/profile-picture", routes.UploadProfilePicture)

	app.Post("admin/users", routes.AddUser)
	app.Post("admin/users/add-csv", routes.AddUsersCsv)
	app.Put("admin/users/:id", routes.EditUser)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
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
