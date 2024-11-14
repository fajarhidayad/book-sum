package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []*fiber.Error {
	var errors []*fiber.Error
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fiber.NewError(fiber.StatusBadRequest, err.StructField()+" is "+err.Tag()))
		}
	}

	return errors
}
