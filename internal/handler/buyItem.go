package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/service"
)

type MerchHandler struct {
	lgr *slog.Logger
	srv *service.MerchService
}

func NewMerchHandler(lgr *slog.Logger, srv *service.MerchService) *MerchHandler {
	return &MerchHandler{lgr: lgr, srv: srv}
}

func (hnd *MerchHandler) BuyItem(ctx *gin.Context) {
	hnd.lgr.Info("Buying item")

	itemName := ctx.Param("item")

	senderID, exists := ctx.Get("user_id")
	if !exists {
		hnd.lgr.Error("Sender dose not exist")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	convertedSenderID, ok := senderID.(float64)
	if !ok {
		hnd.lgr.Error("Error while converting to float")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id"})
		return
	}

	if err := hnd.srv.BuyItem(int(convertedSenderID), itemName); err != nil {
		hnd.lgr.Error("Error while buying item", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while buying item"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Item was purchased successfully"})
}
