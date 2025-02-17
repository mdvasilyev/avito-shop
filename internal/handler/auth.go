package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/model"
	"github.com/mdvasilyev/avito-shop/internal/service"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	lgr *slog.Logger
	srv *service.AuthService
}

func NewAuthHandler(lgr *slog.Logger, srv *service.AuthService) *AuthHandler {
	return &AuthHandler{lgr: lgr, srv: srv}
}

func (hnd *AuthHandler) Auth(ctx *gin.Context) {
	hnd.lgr.Info("Authing")

	var req model.AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		hnd.lgr.Error("Error while binding json", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while binding json"})
		return
	}

	token, err := hnd.srv.Auth(req.Username, req.Password)
	if err != nil {
		hnd.lgr.Error("Error while getting info", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting info"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
