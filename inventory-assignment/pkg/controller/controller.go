package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Alok-Majumder/inventory-assignment/internal/environment"
	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
	"github.com/Alok-Majumder/inventory-assignment/internal/products"
)

type Controller struct {
	Envs  *environment.Environment
	ctx   context.Context
	sqlDB *postgres.PostGres
}

func NewController(envs *environment.Environment, ctx context.Context, postGres *postgres.PostGres) *Controller {
	return &Controller{

		Envs:  envs,
		ctx:   ctx,
		sqlDB: postGres,
	}
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pRepo := products.NewProductRepository(c.sqlDB)
	pServ := products.NewService(pRepo)

	// Check if it is a Get Method and to fetch all Products URI
	if r.URL.Path == "/products/getAll" && r.Method == http.MethodGet {

		prdqty, _ := pServ.GetAllProducts(c.ctx)
		encodeResponseAsJSON(prdqty, w)

	} else if r.URL.Path == "/products/sell" && r.Method == http.MethodPost {
		jsonString, _ := generateJsonString(r)
		res, err := pServ.SellProduct(c.ctx, jsonString)
		if err != nil {
			fmt.Println("Error At the time of selling", err)
		}
		encodeResponseAsJSON(res, w)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		fmt.Println("Error Writing response body", err)
	}
}

func generateJsonString(r *http.Request) ([]byte, error) {
	jsonRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return jsonRequest, nil

}
