package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
)

type InventoryRepository interface {
	GetInventoryByID(ctx context.Context, id string, inventory *Inventory) error
	UpdateInventoryByID(ctx context.Context, id string, inventory *Inventory) error
}

type InventoryRepo struct {
	SqlDB *postgres.PostGres
}

func NewInventoryRepo(postGres *postgres.PostGres) *InventoryRepo {
	return &InventoryRepo{
		SqlDB: postGres,
	}
}

func (r *InventoryRepo) GetInventoryByID(ctx context.Context, id string, inventory *Inventory) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	sqlStatement := `SELECT ART_ID, NAME, STOCK FROM INVENTORY WHERE ART_ID = $1`

	// Will return single row

	err := r.SqlDB.DB.QueryRowContext(ctx, sqlStatement, id).Scan(&inventory.ArtID, &inventory.Name, &inventory.Stock)
	if err != nil {
		fmt.Println("Erro....", err)
		return fmt.Errorf("Rows.Scan: %v", err)
	}
	fmt.Println("  GetInventoryByID", inventory)

	return nil
}
