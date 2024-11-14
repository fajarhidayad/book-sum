package main

import (
	"log"

	"github.com/fajarhidayad/book-sum/config"
	"github.com/fajarhidayad/book-sum/database"
	"github.com/fajarhidayad/book-sum/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type GlobalErrorHandler struct {
	Message string `json:"message"`
}

func init() {
	database.ConnectDB()
}

func main() {
	PORT := config.Config("PORT")

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandler{
				Message: err.Error(),
			})
		},
	})
	app.Use(cors.New())

	router.Routes(app)

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}
