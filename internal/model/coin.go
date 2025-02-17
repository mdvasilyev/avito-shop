package model

type CoinTransactionRequest struct {
	UserID int `json:"userID"`
	Amount int `json:"amount"`
}

type CoinTransaction struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"senderID"`
	ReceiverID int    `json:"receiverID"`
	Quantity   int    `json:"quantity"`
	CreatedAt  string `json:"createdAt"`
}

type CoinUser struct {
	ID    int `json:"id"`
	Coins int `json:"coins"`
}
