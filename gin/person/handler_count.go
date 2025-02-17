package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pDB PersonService) CountPeople(c *gin.Context) {
	num, err := pDB.DB.Count()

	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, num)
}
