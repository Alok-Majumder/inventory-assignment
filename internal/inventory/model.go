package inventory

type Inventory struct {
	ArtID string  `json:"art_id"`
	Name  string  `json:"name"`
	Stock float64 `json:"stock"`
}
