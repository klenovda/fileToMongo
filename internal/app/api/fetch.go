package api

import (
	"context"
	"fileToMongo/pkg/apipb"
	"net/http"

	"github.com/pkg/errors"
)

// Fetch CSV-file with product list
func (i *Implementation) Fetch(ctx context.Context, req *apipb.FetchRequest) (*apipb.FetchResponse, error) {
	if err := i.product.FetchCSV(ctx, req.Url); err != nil {
		return &apipb.FetchResponse{
			Status: http.StatusInternalServerError,
		}, errors.Wrap(err, "couldn't fetch file")
	}

	return &apipb.FetchResponse{
		Status: http.StatusOK,
	}, nil
}
