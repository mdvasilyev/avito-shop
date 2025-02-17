package model

type Merch struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type BuyRequest struct {
	ItemName string `json:"item"`
}

type ErrorResponse struct {
	Error string `json:"errors"`
}
