package dbs

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ssinist3r/lightweight-go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var ctx = context.Background()

func NewMongoDB() *MongoDB {
	var collection *mongo.Collection

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil) // check if connection is ready
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("lightweight-go-server").Collection("people")
	collection.Drop(ctx) // drop any previous document stored

	// load default data
	loadData := make([]interface{}, len(People))
	for i := range People {
		ind, _ := strconv.Atoi(i)
		loadData[ind-1] = People[i]
	}
	if _, err := collection.InsertMany(ctx, loadData); err != nil {
		panic(err)
	}

	// create indexes
	defLanguage := "english"
	defUnique := true
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "id", Value: 1}}, // id in asc order
		}, {
			Keys: bson.D{{Key: "nickname", Value: 1}},
			Options: &options.IndexOptions{
				Unique: &defUnique,
			},
		}, {
			Keys: bson.D{{Key: "name", Value: "text"}, {Key: "nickname", Value: "text"}},
			Options: &options.IndexOptions{
				DefaultLanguage: &defLanguage,
			},
		},
	}

	if _, err := collection.Indexes().CreateMany(ctx, indexes); err != nil {
		panic(fmt.Errorf("Error creating DB indexes: %w", err))
	}

	return &MongoDB{
		client,
		collection,
	}
}

func (db *MongoDB) Create(p *models.Person) error {
	_, err := db.collection.InsertOne(ctx, p)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ErrDuplicateKey
		}

		return fmt.Errorf("Error inserting person: %w", err)
	}

	return nil
}

func (db *MongoDB) Get(id string) (*models.Person, error) {
	p := new(models.Person)

	err := db.collection.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(p)

	switch err {
	case nil:
		return p, nil
	case mongo.ErrNoDocuments:
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("Error decoding get result for term (%s): %w", id, err)
	}
}

func (db *MongoDB) Search(term string) ([]*models.Person, error) {
	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: term}}}}
	opts := options.Find().SetLimit(50)

	cursor, err := db.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("Error searching for term (%s): %w", term, err)
	}
	defer cursor.Close(ctx)

	var results []*models.Person
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("Error decoding search results for term (%s): %w", term, err)
	}

	return results, nil
}

func (db *MongoDB) Count() (int, error) {
	count, err := db.collection.CountDocuments(ctx, bson.D{{}})

	if err != nil {
		return 0, err
	}

	return int(count), nil
}
