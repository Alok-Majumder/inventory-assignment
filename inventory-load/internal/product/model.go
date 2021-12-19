package product

type ProductsSrc struct {
	Products []ProductSrc `json:"products"`
}
type ContainArticles struct {
	ArtID    string `json:"art_id"`
	AmountOf string `json:"amount_of"`
}
type ProductSrc struct {
	Name            string            `json:"name"`
	ContainArticles []ContainArticles `json:"contain_articles"`
}

type Product struct {
	ProductName string
	ArtID       string
	AmountOfArt int
}
