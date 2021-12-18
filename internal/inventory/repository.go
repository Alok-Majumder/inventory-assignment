package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/Alok-Majumder/inventory-assignment/internal/postgres"
)

type InventoryRepository interface {
	GetInventoryByID(ctx context.Context, id string, inventory *Inventory) error
	UpdateInventoryByID(ctx context.Context, inventory *Inventory) error
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
		fmt.Println("Erro....From ... ", err)
		return fmt.Errorf("Rows.Scan: %v", err)
	}

	return nil
}

func (r *InventoryRepo) UpdateInventoryByID(ctx context.Context, id string, reduceby int) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	sqlStatement := `UPDATE INVENTORY SET STOCK = STOCK - $1 WHERE ART_ID = $2`

	stmt, err := r.SqlDB.DB.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reduceby, id)
	if err != nil {
		return err
	}

	return nil
}
