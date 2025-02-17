package dbs

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/ssinist3r/lightweight-go-server/models"
)

type postgresDB struct {
	conn *gorm.DB
}

var loadData = []*models.Person{
	{ID: "1", Name: "SSinist3r", Nickname: "Rosner", Birth: "2000-10-01"},
	{ID: "2", Name: "João Rosner", Nickname: "João", Birth: "2000-10-02"},
	{ID: "3", Name: "Maria", Nickname: "Maria", Birth: "2000-10-03"},
}

func NewPostgresDB() *postgresDB {
	uri := os.Getenv("POSTGRES_URI") // to use with docker compose
	// uri := "postgres://postgres:1234@localhost/lightweight-go-server?sslmode=disable" // to use with docker
	conn, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic(fmt.Errorf("Error connecting to postgres (creds: %s):%q", uri, err))
	}

	fmt.Println("Database is connected")

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}))

	if err != nil {
		panic(fmt.Errorf("Error initializing gorm: %q", err))
	}

	db.Migrator().DropTable(&models.Person{})   // drop any previous data stored
	db.Migrator().CreateTable(&models.Person{}) // create people table
	// create indexes
	// For id field
	db.Migrator().CreateIndex(&models.Person{}, "ID")
	db.Migrator().CreateIndex(&models.Person{}, "idx_id")
	// For Name field
	db.Migrator().CreateIndex(&models.Person{}, "Name")
	db.Migrator().CreateIndex(&models.Person{}, "idx_name")
	// For Nickname field
	db.Migrator().CreateIndex(&models.Person{}, "Nickname")
	db.Migrator().CreateIndex(&models.Person{}, "idx_nickname")
	// For full text search on name and nickname
	db.Exec("CREATE INDEX idx_name_ftext ON people USING GIN (to_tsvector('english', name))")
	db.Exec("CREATE INDEX idx_nickname_ftext ON people USING GIN (to_tsvector('english', nickname))")

	resLoad := db.Create(loadData)
	if resLoad.Error != nil {
		panic(resLoad.Error)
	}

	return &postgresDB{
		conn: db,
	}
}

func (db *postgresDB) Create(p *models.Person) error {
	result := db.conn.Create(p)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrDuplicateKey
		}

		return fmt.Errorf("Error inserting person: %w", result.Error)
	}

	return nil
}

func (db *postgresDB) Get(id string) (*models.Person, error) {
	var person = models.Person{ID: id}
	err := db.conn.First(&person)

	switch err.Error {
	case nil:
		return &person, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("Error decoding get result for term (%s): %s", id, err.Error.Error())
	}
}

func (db *postgresDB) Search(term string) ([]*models.Person, error) {
	var people []*models.Person
	substring := "%" + term + "%"

	err := db.conn.Limit(50).Where("name LIKE ?", substring).Find(&models.Person{}).Scan(&people)

	if err.Error != nil {
		return nil, fmt.Errorf("Error searching for term (%s): %s", term, err.Error.Error())
	}

	return people, nil
}

func (db *postgresDB) Count() (int, error) {
	var count int64
	result := db.conn.Model(&models.Person{}).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}
