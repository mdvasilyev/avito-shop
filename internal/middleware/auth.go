package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/helper"
	"log/slog"
	"net/http"
	"strings"
)

func AuthMiddleware(lgr *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lgr.Info("Applying AuthMiddleware")

		header := ctx.GetHeader("Authorization")
		if header == "" {
			lgr.Error("Missing authorization header")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims, err := helper.ValidateToken(tokenString)
		if err != nil {
			lgr.Error("Invalid token")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", (*claims)["user_id"])

		ctx.Next()
	}
}
