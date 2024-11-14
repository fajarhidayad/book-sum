package tests

import (
	"bytes"
	"encoding/json"
	"github.com/fajarhidayad/book-sum/database"
	"github.com/fajarhidayad/book-sum/model"
	"github.com/fajarhidayad/book-sum/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func init() {
	database.ConnectDB()
}

func setupApp() *fiber.App {
	app := fiber.New()
	router.Routes(app)
	return app
}

func TestCreateAuthor(t *testing.T) {
	app := setupApp()

	var author model.Author
	database.DB.Where("1=1").Delete(&author)

	payload := map[string]string{
		"name": "Henry",
		"bio":  "Best selling author",
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/v1/authors", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var response map[string]string
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "created", response["message"])
}

func TestGetAllAuthors(t *testing.T) {
	app := setupApp()

	var authors []model.Author

	req := httptest.NewRequest("GET", "/api/v1/authors", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	err := json.NewDecoder(resp.Body).Decode(&authors)
	if err != nil {
		panic(err)
	}
}
