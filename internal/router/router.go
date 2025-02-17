package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/handler"
	"github.com/mdvasilyev/avito-shop/internal/middleware"
	"log/slog"
)

func SetupRoutes(lgr *slog.Logger, rtr *gin.Engine, dbConn *sql.DB) {
	lgr.Info("Setting up routes")

	apiGroup := rtr.Group("/api")

	authService := handler.NewAuthService(dbConn)
	merchService := handler.NewMerchService(dbConn)
	userService := handler.NewUserService(dbConn)
	coinService := handler.NewCoinService(dbConn)

	apiGroup.POST("/auth", authService.Auth)
	apiGroup.POST("/sendCoin", middleware.AuthMiddleware(lgr), coinService.SendCoin)
	apiGroup.GET("/info", middleware.AuthMiddleware(lgr), userService.Info)
	apiGroup.GET("/buy/:item", middleware.AuthMiddleware(lgr), merchService.BuyItem)
}
