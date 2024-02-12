// pkg/api/middleware/auth.go

package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware is middleware for JWT token authentication
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        // Validate JWT token here...

        c.Next()
    }
}
