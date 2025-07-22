package middleware

import (
	"github.com/NeginSal/job-tracker-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}