package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logrus logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Initialize native SQL with PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=tech_assess port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database: ", err)
		return
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Register endpoints from routes.go
	registerRoutes(db, router /*, logger */) // Pass logger if needed

	logger.Info("Starting server at :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Server failed: ", err)
	}
}
