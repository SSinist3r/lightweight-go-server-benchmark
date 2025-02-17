package person

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ssinist3r/lightweight-go-server/models"
)

func (pDB PersonService) AddPerson(c echo.Context) error {
	var person models.Person

	if err := c.Bind(&person); err != nil {
		c.JSON(http.StatusBadRequest, "ValidationError: "+err.Error())
	}

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// people = append(people, person)

	if err := pDB.DB.Create(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, string(err.Error()))
	}

	return c.JSON(http.StatusCreated, string(jsonPerson))
}
