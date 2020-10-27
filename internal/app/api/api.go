//go:generate minimock -i Product -o ./mocks/product.go

package api

import (
	"context"
	"fileToMongo/internal/database"
	"fileToMongo/pkg/apipb"
)

type Product interface {
	FetchCSV(ctx context.Context, u string) error
	List(ctx context.Context, page *apipb.ListRequest_PagingParams, sort *apipb.ListRequest_SortingParams) ([]*database.Product, error)
}

// Implementation ...
type Implementation struct {
	product Product
}

func NewApi(product Product) *Implementation {
	return &Implementation{
		product: product,
	}
}
