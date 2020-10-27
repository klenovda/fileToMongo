package api

import (
	"context"
	"fileToMongo/pkg/apipb"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseList struct {
	name    string
	want    *apipb.ListResponse
	wantErr bool
	err     error
}

func TestImplementation_List(t *testing.T) {
	for _, test := range getTestCaseList() {
		t.Run(test.name, func(t *testing.T) {
			i := createImplementation(t, test.err)

			resp, err := i.List(context.Background(), getListRequest())

			if test.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, errors.Wrap(test.err, "couldn't get list of products").Error())
				assert.Equal(t, test.want, resp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.want, resp)
			}
		})
	}
}

func getListRequest() *apipb.ListRequest {
	return &apipb.ListRequest{
		PagingParams: &apipb.ListRequest_PagingParams{
			Page:  1,
			Limit: 10,
		},
		SortingParams: &apipb.ListRequest_SortingParams{
			Param: "Name",
			Sort:  apipb.Sort_ASC,
		},
	}
}

func getTestCaseList() []testCaseList {
	return []testCaseList{
		{
			name: "positive case",
			want: &apipb.ListResponse{
				Product: []*apipb.ListResponse_Product{
					{
						Name:      "product",
						Price:     100,
						CreatedAt: tsNow,
					},
				},
			},
			wantErr: false,
		},
	}
}
