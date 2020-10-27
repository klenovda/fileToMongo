package api

import (
	"context"
	"fileToMongo/internal/database"
	"fileToMongo/pkg/apipb"

	"github.com/pkg/errors"
)

// List get list of products
func (i *Implementation) List(ctx context.Context, req *apipb.ListRequest) (*apipb.ListResponse, error) {
	p, err := i.product.List(ctx, req.PagingParams, req.SortingParams)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of products")
	}

	return convertToListResponse(p), nil
}

func convertToListResponse(products []*database.Product) *apipb.ListResponse {
	listProducts := make([]*apipb.ListResponse_Product, 0, len(products))

	for _, p := range products {
		lp := &apipb.ListResponse_Product{
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
		}

		listProducts = append(listProducts, lp)
	}

	return &apipb.ListResponse{Product: listProducts}
}
