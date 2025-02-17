package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PassHash string `json:"pass_hash"`
	Coins    int    `json:"coins"`
}

type InventoryItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CoinHistory struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

type UserResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}
