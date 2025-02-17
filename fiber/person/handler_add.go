package person

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ssinist3r/lightweight-go-server/models"
)

func (pDB PersonService) AddPerson(c *fiber.Ctx) error {
	var person models.Person

	if err := c.BodyParser(&person); err != nil {
		return c.Status(http.StatusBadRequest).SendString("ValidationError: " + err.Error())
	}

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Bad Request")
	}

	if err := pDB.DB.Create(&person); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(string(err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(string(jsonPerson))
}
