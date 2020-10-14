package database

import (
	"context"
	paginate "github.com/gobeam/mongo-go-pagination"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const collectionName string = "products"

type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
	Qty uint `json:"qty" bson:"qty"`
	CreatedAt *timestamp.Timestamp `json:"created_at" bson:"created_at"`
}

type Storage struct {
	db *mongo.Database
}

func NewStorage(client *mongo.Client, dbName string) *Storage {
	return &Storage{
		db: client.Database(dbName),
	}
}

// Find list of products form storage
func (s Storage) Find(ctx context.Context, limit int64, page int64, sortField string, sortValue int, filters interface{}) ([]*Product, error) {
	collection := s.db.Collection(collectionName)
	paginatedData, err := paginate.New(collection).
		Limit(limit).
		Page(page).
		Sort(sortField, sortValue).
		Filter(filters).
		Find()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find paginatedData")
	}

	var products []*Product
	for _, raw := range paginatedData.Data {
		var product *Product
		if marshallErr := bson.Unmarshal(raw, &product); marshallErr == nil {
			products = append(products, product)
		}

	}

	return products, nil
}

// Insert products to storage
func (s Storage) Insert(ctx context.Context, products map[string]*Product) error {
	collection := s.db.Collection(collectionName)
	for _, p := range products {
		log.Printf("product %v", p)
		res, err := collection.UpdateOne(ctx,
			bson.M{"name": p.Name},
			bson.D{
				{"$set", bson.D{
					{"name", p.Name},
					{"price", p.Price},
					{"qty", p.Qty},
					{"created_at", p.CreatedAt},
				}},
			},
			options.Update().SetUpsert(true))
		if err != nil {
			return errors.Wrapf(err, "couldn't insert product with name %s", p.Name)
		}
		log.Printf("upserted product with id %v", res.UpsertedID)
	}

	return nil
}