package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, exists := ctx.Get("userClaims")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user claims do not exist"})
			return
		}

		claimsMap, ok := userClaims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cannot convert claims to map"})
			return
		}

		userIdString, ok := claimsMap["user_id"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cannot get user_id from map"})
			return
		}

		userId, err := uuid.Parse(userIdString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		// You can store claims in context for use in handlers
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
