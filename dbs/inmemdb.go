package dbs

import "github.com/ssinist3r/lightweight-go-server/models"

type InMemDB struct {
	data map[string]*models.Person
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		data: People,
	}
}

func (db *InMemDB) Create(p *models.Person) error {
	db.data[p.ID] = p
	return nil
}

func (db *InMemDB) Get(id string) (*models.Person, error) {
	p, ok := db.data[id]

	if !ok {
		return nil, ErrNotFound
	}

	return p, nil
}

func (db *InMemDB) Count() (int, error) {
	return len(db.data), nil
}

func (db *InMemDB) Search(term string) ([]*models.Person, error) {
	return make([]*models.Person, 0), nil
}
