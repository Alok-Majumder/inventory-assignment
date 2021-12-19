package inventory

import "github.com/Alok-Majumder/inventory-load/internal/postgres"

type Repository interface {
	SetInventory(model InventorySrc) error
	GetInventory(key string) (InventorySrc, error)
	UpdateInventory(model InventorySrc) error
}

type postgresrepo struct {
	client *postgres.PostGres
}

// NewBigtableRepository - creates and initializes a bigtable repository.
func NewPostGresRepo(postgres *postgres.PostGres) Repository {
	return &postgresrepo{
		postgres,
	}
}

func (p *postgresrepo) SetInventory(model InventorySrc) error {
	//TODO
	return nil
}

func (p *postgresrepo) GetInventory(key string) (InventorySrc, error) {
	//TODO
	return InventorySrc{}, nil
}

func (p *postgresrepo) UpdateInventory(model InventorySrc) error {
	//TODO
	return nil
}
