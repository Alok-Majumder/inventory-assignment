package inventory_test

import (
	"testing"

	"github.com/Alok-Majumder/inventory-load/internal/inventory"
	"github.com/stretchr/testify/assert"
)

type repository struct {
	persistedModel inventory.InventorySrc
}

func (r *repository) SetInventory(model inventory.InventorySrc) error {
	r.persistedModel = model
	return nil

}

func (r *repository) GetInventory(key string) (inventory.InventorySrc, error) {
	r.persistedModel.ArtID = key
	return inventory.InventorySrc{}, nil

}

func (r *repository) UpdateInventory(model inventory.InventorySrc) error {
	r.persistedModel = model
	return nil

}

func TestUpdateWhenExists(t *testing.T) {
	r := &repository{}
	service := inventory.NewService(r)

	inventoriesSrc := new(inventory.InventoriesSrc)

	err := service.Process(inventoriesSrc)
	assert.Error(t, err)

}
