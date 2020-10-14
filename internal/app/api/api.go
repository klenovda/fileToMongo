package api

import (
	"context"
	"testTask/internal/database"
	"testTask/pkg/api"
)

type Product interface {
	FetchCSV(ctx context.Context, u string) error
	List(ctx context.Context, page *api.ListRequest_PagingParams, sort *api.ListRequest_SortingParams) ([]*database.Product, error)
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