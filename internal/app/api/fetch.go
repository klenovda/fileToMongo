package api

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	desc "testTask/pkg/api"
)

// Fetch CSV-file with product list
func (i *Implementation) Fetch(ctx context.Context, req *desc.FetchRequest) (*desc.FetchResponse, error) {
	if err := i.product.FetchCSV(ctx, req.Url); err != nil {
		return &desc.FetchResponse{
			Status: http.StatusInternalServerError,
		}, errors.Wrap(err, "couldn't fetch file")
	}

	return &desc.FetchResponse{
		Status: http.StatusOK,
	}, nil
}