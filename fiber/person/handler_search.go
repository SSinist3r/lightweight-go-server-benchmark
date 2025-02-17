package person

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ssinist3r/lightweight-go-server/dbs"
)

func (pDB PersonService) SearchPerson(c *fiber.Ctx) error {
	term := c.Query("t")

	if term == "" {
		return fiber.ErrBadRequest
	}

	p, err := pDB.DB.Search(term)

	switch err {
	case nil:
		return c.Status(http.StatusOK).JSON(p)
	case dbs.ErrNotFound:
		return fiber.ErrNotFound
	default:
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
}
