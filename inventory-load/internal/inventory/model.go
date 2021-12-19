package inventory

type InventoriesSrc struct {
	Inventory []InventorySrc `json:"inventory"`
}
type InventorySrc struct {
	ArtID string `json:"art_id"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
}

type Inventory struct {
	ART_ID string
	NAME   string
	STOCK  int
}
