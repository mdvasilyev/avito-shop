package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/model"
	"github.com/mdvasilyev/avito-shop/internal/service"
)

type CoinHandler struct {
	lgr *slog.Logger
	srv *service.CoinService
}

func NewCoinHandler(lgr *slog.Logger, srv *service.CoinService) *CoinHandler {
	return &CoinHandler{lgr: lgr, srv: srv}
}

func (hnd *CoinHandler) SendCoin(ctx *gin.Context) {
	hnd.lgr.Info("Sending coins")

	var req model.CoinTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		hnd.lgr.Error("Error while binding json", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while binding json"})
		return
	}

	senderID, exists := ctx.Get("user_id")
	if !exists {
		hnd.lgr.Error("Sender does not exist")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	convertedSenderID, ok := senderID.(float64)
	if !ok {
		hnd.lgr.Error("Error while converting to float")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id"})
		return
	}

	if err := hnd.srv.SendCoins(int(convertedSenderID), req.UserID, req.Amount); err != nil {
		hnd.lgr.Error("Error while sending coins", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while sending coins"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Coin transaction is successful"})
}
