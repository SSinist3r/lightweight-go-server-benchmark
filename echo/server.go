package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"github.com/ssinist3r/lightweight-go-server/dbs"
	"github.com/ssinist3r/lightweight-go-server/echo/person"
)

func main() {
	e := echo.New()

	var db dbs.DB

	switch os.Getenv("DB") {
	case "mongo":
		println("Setting mongoDB")
		db = dbs.NewMongoDB()
	case "postgres":
		println("Setting PostgresDB")
		db = dbs.NewPostgresDB()
	default:
		println("Setting InMemoryDB")
		db = dbs.NewInMemDB()
	}

	service := person.PersonService{DB: db}

	e.POST("/people", service.AddPerson)
	e.GET("/people", service.SearchPerson)
	e.GET("/people/:id", service.GetPerson)
	e.GET("/contagem-people", service.CountPeople)

	e.Logger.Fatal(e.Start(":8080"))
}
