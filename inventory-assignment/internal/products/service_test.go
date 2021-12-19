package products_test

import (
	"context"
	"testing"

	"github.com/Alok-Majumder/inventory-assignment/internal/products"
	"github.com/stretchr/testify/assert"
)

type testRepository struct {
}

func TestGetAllProductsWithQuantity(t *testing.T) {

	ctx := context.Background()
	repo := &testRepository{}

	service := products.NewService(repo)

	_, err := service.GetAllProducts(ctx)
	// Assert
	assert.NoError(t, err)
}

func (r *testRepository) GetAllProductsWithQuantity(ctx context.Context) ([]*products.ProductQty, error) {
	return nil, nil
}

func (r *testRepository) GetProductByName(ctx context.Context, key string) ([]*products.Product, error) {
	return nil, nil
}

func (r *testRepository) SellProduct(ctx context.Context, product []byte) (*products.Response, error) {
	return nil, nil
}

func (r *testRepository) GetProductWithQuantity(ctx context.Context, product products.ProductQty) (*products.ProductQty, error) {
	return nil, nil

}
