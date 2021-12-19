package inventory_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/Alok-Majumder/inventory-assignment/internal/inventory"
	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var NewInventoryMock = &inventory.Inventory{
	ArtID: "1",
	Name:  "leg",
	Stock: 13,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func TestGetInventoryByID(t *testing.T) {
	db, mock := NewMock()
	SqlDB := &postgres.PostGres{
		DB: db,
	}

	defer func() {
		db.Close()
	}()

	r := &inventory.InventoryRepo{
		SqlDB: SqlDB,
	}
	sqlStatement := "SELECT ART_ID, NAME, STOCK FROM INVENTORY WHERE ART_ID = \\$1"

	rows := sqlmock.NewRows([]string{"ART_ID", "NAME", "STOCK"}).
		AddRow(NewInventoryMock.ArtID, NewInventoryMock.Name, NewInventoryMock.Stock)

	mock.ExpectQuery(sqlStatement).WithArgs(NewInventoryMock.ArtID).WillReturnRows(rows)
	ctx := context.Background()
	inventory := new(inventory.Inventory)

	err := r.GetInventoryByID(ctx, NewInventoryMock.ArtID, inventory)
	fmt.Println(inventory)
	assert.NotNil(t, inventory)
	assert.NoError(t, err)

}
