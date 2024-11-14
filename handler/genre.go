package handler

import (
	"fmt"

	"github.com/fajarhidayad/book-sum/database"
	"github.com/fajarhidayad/book-sum/model"
	"github.com/fajarhidayad/book-sum/validation"
	"github.com/gofiber/fiber/v2"
)

func GetAllGenres(c *fiber.Ctx) error {
	db := database.DB
	var genres []model.Genre

	db.Find(&genres).Limit(10)

	return c.JSON(fiber.Map{
		"data": genres,
	})
}

func GetGenreById(c *fiber.Ctx) error {
	db := database.DB
	var genre model.Genre
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id type",
		})
	}

	db.First(&genre, id)

	if genre.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "genre not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": genre,
	})
}

func AddGenre(c *fiber.Ctx) error {
	db := database.DB
	var genre model.Genre

	if err := c.BodyParser(&genre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid input",
			"message": err,
		})
	}

	errors := validation.ValidateStruct(genre)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	result := db.Create(&genre)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateGenre(c *fiber.Ctx) error {
	db := database.DB
	var genre model.Genre
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id type",
		})
	}

	db.First(&genre, id)
	if genre.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "genre not found",
		})
	}

	if err := c.BodyParser(&genre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errors := validation.ValidateStruct(genre)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db.Model(&genre).Where("id = ?", id).Updates(&genre)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteGenreById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Invalid type of id",
		})
	}

	db := database.DB
	var genre model.Genre
	db.First(&genre, id)

	if genre.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("genre with id %d not found", id),
		})
	}

	db.Delete(&genre)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
