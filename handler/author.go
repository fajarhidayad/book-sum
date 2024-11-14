package handler

import (
	"fmt"

	"github.com/fajarhidayad/book-sum/database"
	"github.com/fajarhidayad/book-sum/model"
	"github.com/fajarhidayad/book-sum/validation"
	"github.com/gofiber/fiber/v2"
)

func GetAllAuthors(c *fiber.Ctx) error {
	db := database.DB
	var authors []model.Author
	db.Find(&authors).Limit(10)

	return c.JSON(fiber.Map{
		"data": authors,
	})
}

func GetAuthorById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "id type is not valid",
		})
	}

	db := database.DB
	var data model.Author
	db.Find(&data, id)

	if data.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Author with id %d not found", id),
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func AddNewAuthor(c *fiber.Ctx) error {
	db := database.DB
	var author model.Author

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid input",
			"message": err.Error(),
		})
	}

	errors := validation.ValidateStruct(author)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result := db.Create(&author)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
	})
}

func UpdateAuthorById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.DB
	var author model.Author

	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Invalid id type",
		})
	}
	var result model.Author
	db.First(&result, id)
	if result.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "author not found",
		})
	}

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	errors := validation.ValidateStruct(&author)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db.Model(&result).Updates(&author)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteAuthorById(c *fiber.Ctx) error {
	db := database.DB
	var author model.Author
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id type",
		})
	}

	db.Find(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "data not found",
		})
	}

	db.Delete(&author)

	return c.JSON(fiber.Map{
		"message": "success delete",
	})
}
