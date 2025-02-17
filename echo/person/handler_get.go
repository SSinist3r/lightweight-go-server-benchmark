package person

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pDB PersonService) GetPerson(c echo.Context) error {
	p, err := pDB.DB.Get(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, "Person not found")
	}

	return c.JSON(http.StatusOK, p)
}
