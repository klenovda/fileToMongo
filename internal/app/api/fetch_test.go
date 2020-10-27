package api

import (
	"context"
	"fileToMongo/pkg/apipb"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseFetch struct {
	name    string
	want    *apipb.FetchResponse
	wantErr bool
	err     error
}

func TestImplementation_Fetch(t *testing.T) {
	for _, test := range getTestCaseFetch() {
		t.Run(test.name, func(t *testing.T) {
			i := createImplementation(t, test.err)

			resp, err := i.Fetch(context.Background(), getFetchRequest())

			if test.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, errors.Wrap(test.err, "couldn't fetch file").Error())
				assert.Equal(t, test.want, resp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.want, resp)
			}
		})
	}
}

func getFetchRequest() *apipb.FetchRequest {
	return &apipb.FetchRequest{
		Url: "http://test.ru",
	}
}

func getTestCaseFetch() []testCaseFetch {
	return []testCaseFetch{
		{
			name:    "positive case",
			want:    &apipb.FetchResponse{Status: http.StatusOK},
			wantErr: false,
		},
		{
			name:    "negative case",
			want:    &apipb.FetchResponse{Status: http.StatusInternalServerError},
			wantErr: true,
			err:     assert.AnError,
		},
	}
}
