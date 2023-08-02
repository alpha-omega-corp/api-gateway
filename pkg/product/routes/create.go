package routes

import (
	"context"
	"encoding/json"
	"github.com/alpha-omega-corp/api-gateway/pkg/product/pb"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, req bunrouter.Request, s pb.ProductServiceClient) error {
	data := new(CreateRequestBody)
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return err
	}

	res, err := s.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
	})

	if err != nil {
		return err
	}

	return bunrouter.JSON(w, res)
}

type CreateRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}
