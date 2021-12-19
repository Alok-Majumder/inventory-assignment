package inventory

import "fmt"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) Process(event *InventoriesSrc) error {

	for _, i := range event.Inventory {
		inv, err := s.repository.GetInventory(i.ArtID)
		if err != nil {
			fmt.Println("Error getting Inv..", err)
		}
		if inv.ArtID == "" {
			s.repository.SetInventory(i)

		} else {
			s.repository.UpdateInventory(i)

		}

	}

	return nil

}
