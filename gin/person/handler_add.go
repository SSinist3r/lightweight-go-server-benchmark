package person

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssinist3r/lightweight-go-server/models"
)

func (pDB PersonService) AddPerson(c *gin.Context) {
	var person models.Person

	if err := c.Bind(&person); err != nil {
		c.JSON(http.StatusBadRequest, "ValidationError: "+err.Error())
	}

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
	}

	if err := pDB.DB.Create(&person); err != nil {
		c.JSON(http.StatusInternalServerError, string(err.Error()))
	}

	c.JSON(http.StatusCreated, string(jsonPerson))
}
