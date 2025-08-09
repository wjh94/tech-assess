package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var jwtSecret = []byte("your-secret-key") // Replace with your secret

func jwtOptionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err == nil && token.Valid {
				c.Set("jwt", token)
			}
		}
		c.Next()
	}
}

func jwtRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("jwt", token)
		c.Next()
	}
}

func handler(c *gin.Context) {
	c.String(200, "Hello, World!")
}

func profileHandler(c *gin.Context) {
	token, exists := c.Get("jwt")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT required"})
		return
	}
	claims, ok := token.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
		return
	}
	c.JSON(200, gin.H{"user": claims})
}

func registerRoutes(db *sql.DB, router *gin.Engine) {
	// Optional JWT for all routes
	router.Use(jwtOptionalMiddleware())

	router.GET("/", handler)
	// Example protected endpoint
	router.GET("/profile", jwtRequiredMiddleware(), profileHandler)
	// Add more endpoints here as needed, using db and router
}
