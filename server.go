package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ssinist3r/lightweight-go-server/echo/person"

	"github.com/ssinist3r/lightweight-go-server/dbs"
)

func main() {
	e := echo.New()

	e.GET("/", HelloWorld)

	// db := dbs.NewInMemDB()
	// db := dbs.NewMongoDB()
	db := dbs.NewPostgresDB()

	service := person.PersonService{DB: db}

	e.POST("/people", service.AddPerson)
	e.GET("/people", service.SearchPerson)
	e.GET("/people/:id", service.GetPerson)
	e.GET("/contagem-people", service.CountPeople)

	e.Logger.Fatal(e.Start(":8080"))
}
