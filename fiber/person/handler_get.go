package person

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (pDB PersonService) GetPerson(c *fiber.Ctx) error {
	p, err := pDB.DB.Get(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Person not found")
	}

	return c.Status(http.StatusOK).JSON(p)
}
