package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/service"
)

type UserHandler struct {
	lgr *slog.Logger
	srv *service.UserService
}

func NewUserHandler(lgr *slog.Logger, srv *service.UserService) *UserHandler {
	return &UserHandler{lgr: lgr, srv: srv}
}

func (hnd *UserHandler) Info(ctx *gin.Context) {
	hnd.lgr.Info("Getting info")

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

	info, err := hnd.srv.Info(int(convertedSenderID))
	if err != nil {
		hnd.lgr.Error("Error while getting info", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting info"})
		return
	}

	ctx.JSON(http.StatusOK, info)
}
