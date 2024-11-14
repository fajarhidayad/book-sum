package router

import (
	"github.com/fajarhidayad/book-sum/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", logger.New())

	// Author routes
	author := v1.Group("/authors")
	author.Get("/", handler.GetAllAuthors)
	author.Get("/:id", handler.GetAuthorById)
	author.Post("/", handler.AddNewAuthor)
	author.Put("/:id", handler.UpdateAuthorById)
	author.Delete("/:id", handler.DeleteAuthorById)

	// Genre routes
	genre := v1.Group("/genres")
	genre.Get("/", handler.GetAllGenres)
	genre.Get("/:id", handler.GetGenreById)
	genre.Post("/", handler.AddGenre)
	genre.Put("/:id", handler.UpdateGenre)
	genre.Delete("/:id", handler.DeleteGenreById)

	// Book routes
	book := v1.Group("/books")
	book.Get("/", handler.GetAllBooks)
	book.Get("/:id", handler.GetBookById)
}
