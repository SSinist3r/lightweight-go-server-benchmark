package person

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (pDB PersonService) CountPeople(c *fiber.Ctx) error {
	num, err := pDB.DB.Count()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(num)
}
