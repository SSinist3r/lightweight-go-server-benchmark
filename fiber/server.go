package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/ssinist3r/lightweight-go-server/dbs"
	"github.com/ssinist3r/lightweight-go-server/fiber/person"
)

func main() {
	f := fiber.New()

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

	f.Post("/people", service.AddPerson)
	f.Get("/people", service.SearchPerson)
	f.Get("/people/:id", service.GetPerson)
	f.Get("/contagem-people", service.CountPeople)

	f.Listen(":8080")
}
