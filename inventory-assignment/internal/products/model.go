package products

type Product struct {
	ProductName string
	ArtID       string
	AmountOfArt int
}

type ProductQty struct {
	ProductName string `json:"ProductName"`
	Quantity    int    `json:"Quantity"`
}

type Response struct {
	ProductName       string `json:"ProductName"`
	RequestedQuantity int    `json:"RequestedQuantity"`
	AvailableQuantity int    `json:"AvailableQuantity"`
	StausMessage      string `json:"StausMessage"`
}
