package products

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Alok-Majumder/inventory-assignment/internal/inventory"
	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
)

type ProductRepository interface {
	GetProductByName(ctx context.Context, key string) ([]*Product, error)
	GetAllProductsWithQuantity(ctx context.Context) ([]*ProductQty, error)
	GetProductWithQuantity(ctx context.Context, inProduct ProductQty) (*ProductQty, error)
	SellProduct(ctx context.Context, product []byte) (*Response, error)
}

type ProductRepo struct {
	SqlDB *postgres.PostGres
}

func NewProductRepository(postgres *postgres.PostGres) ProductRepository {
	return &ProductRepo{
		SqlDB: postgres,
	}

}

// Provide list of rows for a product with the artid and their required amount to construct the product
func (p *ProductRepo) GetProductByName(ctx context.Context, key string) ([]*Product, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	products := make([]*Product, 0)

	defer cancel()
	sqlStatement := `SELECT PRODUCT_NAME, ART_ID, ART_AMOUNT FROM PRODUCTS WHERE PRODUCT_NAME = $1`

	rows, err := p.SqlDB.DB.QueryContext(ctx, sqlStatement, key)
	if err != nil {
		fmt.Println("Erro....From ... ", err)
		return nil, fmt.Errorf("Rows.Scan: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		product := new(Product)
		err = rows.Scan(&product.ProductName, &product.ArtID, &product.AmountOfArt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)

	}

	return products, nil

}

//Get All Products and their Available Quantity
func (p *ProductRepo) GetAllProductsWithQuantity(ctx context.Context) ([]*ProductQty, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	i := inventory.NewInventoryRepo(p.SqlDB)
	//products := make([]*Product, 0)
	productQtys := make([]*ProductQty, 0)

	defer cancel()
	distinctproduct := `SELECT DISTINCT PRODUCT_NAME FROM PRODUCTS`

	//1. Find Distinct Product from PRODUCTS table
	drows, err := p.SqlDB.DB.QueryContext(ctx, distinctproduct)
	if err != nil {
		fmt.Println("Erro....From ... ", err)
		return nil, fmt.Errorf("Rows.Scan: %v", err)
	}
	defer drows.Close()

	//2. Iterate through the Product List and get product and inventory relation
	for drows.Next() {
		var distProduct string
		err := drows.Scan(&distProduct)
		if err != nil {
			fmt.Println("Error scaning rows ", err)

			return nil, err
		}

		products, err := p.GetProductByName(ctx, distProduct)
		if err != nil {
			return nil, err
		}

		minQty := findAvlQty(products, i, ctx)
		productQty := &ProductQty{
			ProductName: distProduct,
			Quantity:    minQty,
		}

		productQtys = append(productQtys, productQty)

	}

	return productQtys, nil
}

// Sell the product if requested amount is avaible
func (p *ProductRepo) SellProduct(ctx context.Context, inProductQty []byte) (*Response, error) {
	// Unmarshal the Product
	productQty := ProductQty{}

	err := json.Unmarshal(inProductQty, &productQty)
	if err != nil {
		fmt.Println("Unmarshal Error ", err)
		return nil, err
	}

	minAvlProd, err := p.GetProductWithQuantity(ctx, productQty)
	if err != nil {
		return nil, err
	}

	if productQty.Quantity > minAvlProd.Quantity {

		response := &Response{
			ProductName:       productQty.ProductName,
			RequestedQuantity: productQty.Quantity,
			AvailableQuantity: minAvlProd.Quantity,
			StausMessage:      "Requested Quentity is greater than The Available Stock",
		}
		return response, nil
	}

	// Fetch the Inventoy required for the product
	products, err := p.GetProductByName(ctx, productQty.ProductName)
	if err != nil {
		fmt.Println("Error Getting Data from Product Table in  SellProduct ", err)
		return nil, err

	}
	i := inventory.NewInventoryRepo(p.SqlDB)

	// Reduce the numbner of Articles from Inventory
	for _, product := range products {
		inventory := new(inventory.Inventory)
		err := i.GetInventoryByID(ctx, product.ArtID, inventory)
		if err != nil {
			fmt.Println("Error from GetInventoryByID ", err)
			return nil, err

		}
		err = i.UpdateInventoryByID(ctx, product.ArtID, product.AmountOfArt*productQty.Quantity)
		if err != nil {
			fmt.Println("Error from UpdateInventoryByID ", err)
			return nil, err

		}

	}

	response := &Response{
		ProductName:       productQty.ProductName,
		RequestedQuantity: productQty.Quantity,
		AvailableQuantity: minAvlProd.Quantity,
		StausMessage:      "Order Has Been Successfully Placed..",
	}

	return response, nil
}

//Get Product and Available Quantity for a given Product
func (p *ProductRepo) GetProductWithQuantity(ctx context.Context, inProduct ProductQty) (*ProductQty, error) {

	products, err := p.GetProductByName(ctx, inProduct.ProductName)
	if err != nil {
		return nil, err
	}

	i := inventory.NewInventoryRepo(p.SqlDB)

	minAvlQty := findAvlQty(products, i, ctx)
	productQty := &ProductQty{
		ProductName: inProduct.ProductName,
		Quantity:    minAvlQty,
	}
	return productQty, nil
}

func findAvlQty(products []*Product, i *inventory.InventoryRepo, ctx context.Context) int {
	prodLength := 0
	minQty := 0

	for _, product := range products {
		inventory := new(inventory.Inventory)
		prodLength++
		err := i.GetInventoryByID(ctx, product.ArtID, inventory)
		if err != nil {
			fmt.Println("Error from GetInventoryByID ", err)
			return 0

		}

		// FInd Minimum Quentity
		Qty := inventory.Stock / product.AmountOfArt
		if prodLength > 0 && Qty < minQty {
			minQty = Qty

		} else {
			minQty = inventory.Stock / product.AmountOfArt
		}

	}

	return minQty
}
