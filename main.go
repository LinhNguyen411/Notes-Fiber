package main

import (
	"github.com/LinhNguyen411/notes-api-fiber/database"
	"github.com/LinhNguyen411/notes-api-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
