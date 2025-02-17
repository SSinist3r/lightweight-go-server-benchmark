package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pDB PersonService) GetPerson(c *gin.Context) {
	p, err := pDB.DB.Get(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, "Person not found")
	}

	c.JSON(http.StatusOK, p)
}
