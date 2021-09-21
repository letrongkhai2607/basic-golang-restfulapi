package main

import (
	"example.com/m/v2/database"
	"example.com/m/v2/router"
	"github.com/gofiber/template/html"
	"log"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.ConnectDB()
	
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
	defer database.DB.Close()
}