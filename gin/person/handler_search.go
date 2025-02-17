package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssinist3r/lightweight-go-server/dbs"
)

func (pDB PersonService) SearchPerson(c *gin.Context) {
	term := c.Query("t")

	if term == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	p, err := pDB.DB.Search(term)

	switch err {
	case nil:
		c.JSON(http.StatusOK, p)
	case dbs.ErrNotFound:
		c.JSON(http.StatusNotFound, err.Error())
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
