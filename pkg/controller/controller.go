package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Alok-Majumder/inventory-assignment/internal/environment"
	"github.com/Alok-Majumder/inventory-assignment/internal/inventory"
	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
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

	// Check if it is a Get Method and to fetch all Products URI
	if r.URL.Path == "/products/getAll" && r.Method == http.MethodGet {
		fmt.Println("Get All Product")
		i := inventory.NewInventoryRepo(c.sqlDB)
		inventory := new(inventory.Inventory)
		i.GetInventoryByID(c.ctx, "1", inventory)
		fmt.Println(inventory)

	} else if r.URL.Path == "/products/sell" && r.Method == http.MethodPost {

		fmt.Println("Sell the Product")

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

}
