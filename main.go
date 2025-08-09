package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize native SQL with PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=tech_assess port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Register endpoints from routes.go
	registerRoutes(db, router)

	fmt.Println("Starting server at :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Server failed:", err)
	}
}
