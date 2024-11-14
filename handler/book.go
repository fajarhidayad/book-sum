package handler

import (
	"github.com/fajarhidayad/book-sum/database"
	"github.com/fajarhidayad/book-sum/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	var db = database.DB
	var books []model.Book

	db.Find(&books)

	return c.JSON(fiber.Map{
		"data": books,
	})
}

func GetBookById(c *fiber.Ctx) error {
	var db = database.DB
	var book model.Book
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid type of id",
		})
	}

	db.Find(book, id)

	if book.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "book not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": book,
	})
}
