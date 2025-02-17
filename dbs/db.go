package dbs

import (
	"fmt"

	"github.com/ssinist3r/lightweight-go-server/models"
)

var (
	ErrNotFound     = fmt.Errorf("Not found")
	ErrDuplicateKey = fmt.Errorf("Duplicate key")
)

type DB interface {
	Create(*models.Person) error
	Get(string) (*models.Person, error)
	Count() (int, error)
	Search(string) ([]*models.Person, error)
}

var People = map[string]*models.Person{
	"1": {ID: "1", Name: "SSinist3r", Nickname: "Rosner", Birth: "2000-10-01"},
	"2": {ID: "2", Name: "João Rosner", Nickname: "João", Birth: "2000-10-02"},
	"3": {ID: "3", Name: "Maria", Nickname: "Maria", Birth: "2000-10-03"},
}
