package main

import (
	"GoogleAuthv2.0/controllers"
	"GoogleAuthv2.0/initializers"
	middlewares "GoogleAuthv2.0/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"log"
)

func init() {
	initializers.InitializeOAuthConfig()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	engine := html.New("templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//logger
	app.Use(logger.New())
	//routing
	app.Get("/", controllers.Home)
	app.Get("/login", controllers.Login)
	app.Get("/callback", controllers.Callback)
	//protected endpoint
	app.Get("/user", middlewares.RequireAuth, controllers.Profile)

	log.Fatal(app.Listen(":8080"))
}
