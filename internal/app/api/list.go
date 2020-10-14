package api

import (
	"context"
	"github.com/pkg/errors"
	"testTask/internal/database"
	desc "testTask/pkg/api"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	p, err := i.product.List(ctx, req.PagingParams, req.SortingParams)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of products")
	}

	return convertToListResponse(p), nil
}

func convertToListResponse(products []*database.Product) *desc.ListResponse {
	listProducts := make([]*desc.ListResponse_Product, 0, len(products))

	for _, p := range products {
		lp := &desc.ListResponse_Product{
			Name: p.Name,
			Price: p.Price,
			CreatedAt: p.CreatedAt,
		}

		listProducts = append(listProducts, lp)
	}

	return &desc.ListResponse{Product: listProducts}
}