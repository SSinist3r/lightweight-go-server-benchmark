package person

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pDB PersonService) CountPeople(c echo.Context) error {
	num, err := pDB.DB.Count()

	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, num)
}
