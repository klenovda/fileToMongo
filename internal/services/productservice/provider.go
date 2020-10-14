package productservice

import (
	"context"
	"encoding/csv"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"testTask/internal/database"
	desc "testTask/pkg/api"
	"time"
)

const (
	csvExtension string = ".csv"
	csvComma rune = ';'
)

type Storage interface {
	Find(ctx context.Context, limit int64, page int64, sortField string, sortValue int, filters interface{}) ([]*database.Product, error)
	Insert(ctx context.Context, products map[string]*database.Product) error
}

// Product
type Provider struct{
	db Storage
}

// NewProvider create new provider
func NewProvider(db Storage) *Provider {
	return &Provider{
		db: db,
	}
}

func (p *Provider) FetchCSV(ctx context.Context, u string) error {
	if filepath.Ext(u) != csvExtension {
		return nil
	}

	resp, err := http.Get(u)
	if err != nil {
		return errors.Wrapf(err, "couldn't get response by url %s", u)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := getCSVData(resp.Body)
	if err != nil {
		return err
	}

//	in := `name1;100
//name2;200
//name1;150
//name;175
//name1;75
//`
//	data, err := getCSVData(strings.NewReader(in))
//	if err != nil {
//		return err
//	}

	products := make(map[string]*database.Product, len(data))
	if err := fillProducts(products, data); err != nil {
		return err
	}
	log.Printf("result map %v", products)

	if err := p.db.Insert(ctx, products); err != nil {
		return err
	}

	return nil
}

func (p *Provider) List(ctx context.Context, page *desc.ListRequest_PagingParams, sort *desc.ListRequest_SortingParams) ([]*database.Product, error) {
	return p.db.Find(ctx, page.Limit, page.Page, sort.Param, sortValue(sort.Sort), bson.M{})
}

func getResponseByURL(u string) (io.ReadCloser, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get response by url %s", u)
	}

	return resp.Body, nil
}

func getCSVData(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	r.Comma = csvComma
	data, err := r.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read from file")
	}

	return data, nil
}

func sortValue(sort desc.Sort) int {
	if sort == desc.Sort_ASC {
		return 1
	}
	return -1
}

func fillProducts(products map[string]*database.Product, data [][]string) error {
	for _, s := range data {
		if len(s) < 2 {
			continue
		}

		name := s[0]
		price, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return errors.Wrap(err, "couldn't parse price")
		}

		if p, ok := products[name]; ok {
			products[name] = &database.Product{
				Name: name,
				Price: price,
				Qty: p.Qty+1,
				CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
			}
			continue
		}

		products[name] = &database.Product{
			Name: name,
			Price: price,
			Qty: 1,
			CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		}
	}

	return nil
}