package middlewares

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
// 			return
// 		}

// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
// 			return
// 		}

// 		tokenString := parts[1]
// 		claims, err := validateToken(tokenString)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// You can store claims in context for use in handlers
// 		c.Set("userClaims", claims)
// 		c.Next()
// 	}
// }

// func validateToken(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}
// 		return jwtSecret, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		return claims, nil
// 	}

// 	return nil, jwt.ErrTokenMalformed
// }