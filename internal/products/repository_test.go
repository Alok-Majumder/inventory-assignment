package products_test

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"

	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
	"github.com/Alok-Majumder/inventory-assignment/internal/products"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMockDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var NewProductsMock1 = &products.Product{
	ProductName: "Dining Chair",
	ArtID:       "1",
	AmountOfArt: 4,
}
var NewProductsMock2 = &products.Product{
	ProductName: "Dining Chair",
	ArtID:       "2",
	AmountOfArt: 8,
}

func TestProductByName(t *testing.T) {
	db, mock := NewMockDB()

	SqlDB := &postgres.PostGres{
		DB: db,
	}

	defer func() {
		db.Close()
	}()

	r := &products.ProductRepo{
		SqlDB: SqlDB,
	}

	sqlStatement := "SELECT PRODUCT_NAME, ART_ID, ART_AMOUNT FROM PRODUCTS WHERE PRODUCT_NAME = \\$1"

	rows := sqlmock.NewRows([]string{"PRODUCT_NAME", "ART_ID", "ART_AMOUNT"}).
		AddRow(NewProductsMock1.ProductName, NewProductsMock1.ArtID, NewProductsMock1.AmountOfArt).
		AddRow(NewProductsMock2.ProductName, NewProductsMock2.ArtID, NewProductsMock2.AmountOfArt)

	mock.ExpectQuery(sqlStatement).WillReturnRows(rows)
	ctx := context.Background()

	productlist, err := r.GetProductByName(ctx, NewProductsMock1.ProductName)
	assert.NotEmpty(t, productlist)
	assert.NoError(t, err)
	assert.Len(t, productlist, 2)

}

func TestProductRepo_GetAllProductsWithQuantity(t *testing.T) {
	type fields struct {
		SqlDB *postgres.PostGres
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*products.ProductQty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &products.ProductRepo{
				SqlDB: tt.fields.SqlDB,
			}
			got, err := p.GetAllProductsWithQuantity(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepo.GetAllProductsWithQuantity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepo.GetAllProductsWithQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}
