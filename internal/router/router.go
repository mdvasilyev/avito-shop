package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/handler"
	"github.com/mdvasilyev/avito-shop/internal/middleware"
	"github.com/mdvasilyev/avito-shop/internal/repository"
	"github.com/mdvasilyev/avito-shop/internal/service"
	"log/slog"
)

func SetupRoutes(lgr *slog.Logger, rtr *gin.Engine, dbConn *sql.DB) {
	lgr.Info("Setting up routes")

	apiGroup := rtr.Group("/api")

	authHandler := handler.NewAuthHandler(lgr, service.NewAuthService(lgr, repository.NewAuthRepository(lgr, dbConn)))
	merchHandler := handler.NewMerchHandler(lgr, service.NewMerchService(lgr, repository.NewMerchRepository(lgr, dbConn)))
	userHandler := handler.NewUserHandler(lgr, service.NewUserService(lgr, repository.NewUserRepository(lgr, dbConn)))
	coinHandler := handler.NewCoinHandler(lgr, service.NewCoinService(lgr, repository.NewCoinRepository(lgr, dbConn)))

	apiGroup.POST("/auth", authHandler.Auth)
	apiGroup.POST("/sendCoin", middleware.AuthMiddleware(lgr), coinHandler.SendCoin)
	apiGroup.GET("/info", middleware.AuthMiddleware(lgr), userHandler.Info)
	apiGroup.GET("/buy/:item", middleware.AuthMiddleware(lgr), merchHandler.BuyItem)
}
