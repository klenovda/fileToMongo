package api

import (
	"context"
	"fileToMongo/internal/app/api/mocks"
	"fileToMongo/internal/database"
	"fileToMongo/pkg/apipb"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

var tsNow = &timestamp.Timestamp{Seconds: time.Now().Unix()}

func createImplementation(t *testing.T, errIn error) *Implementation {
	product := mocks.NewProductMock(t)
	product.FetchCSVMock.Set(func(ctx context.Context, u string) (err error) {
		return errIn
	})

	retList := []*database.Product{
		{
			Name:      "product",
			Price:     100,
			CreatedAt: tsNow,
		},
	}
	product.ListMock.Set(func(ctx context.Context, page *apipb.ListRequest_PagingParams, sort *apipb.ListRequest_SortingParams) (ppa1 []*database.Product, err error) {
		return retList, errIn
	})

	return NewApi(product)
}
