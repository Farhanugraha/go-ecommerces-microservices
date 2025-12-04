package main

import (
	"auth-service/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User struct untuk response
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Connect to database
	database.Connect()

	// Create Gin router
	r := gin.Default()

	// Route: Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Auth Service API",
			"version": "1.0.0",
			"author":  "Farhan",
		})
	})

	// Route: Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	// API Group v1
	api := r.Group("/api/v1")
	{
		// GET all users
		api.GET("/users", getUsers)

		// GET user by ID
		api.GET("/users/:id", getUserByID)

		// GET user by query parameter (?name=farhan)
		api.GET("/search", searchUser)
	}

	// Start server
	log.Println("Auth Service is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Handler: Get all users
func getUsers(c *gin.Context) {
	// Dummy data (nanti dari database)
	users := []User{
		{ID: 1, Name: "Farhan", Email: "farhan@example.com"},
		{ID: 2, Name: "John Doe", Email: "john@example.com"},
		{ID: 3, Name: "Jane Smith", Email: "jane@example.com"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Users retrieved successfully",
		"data":    users,
		"count":   len(users),
	})
}

// Handler: Get user by ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Dummy data
	users := map[string]User{
		"1": {ID: 1, Name: "Farhan", Email: "farhan@example.com"},
		"2": {ID: 2, Name: "John Doe", Email: "john@example.com"},
		"3": {ID: 3, Name: "Jane Smith", Email: "jane@example.com"},
	}

	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User retrieved successfully",
		"data":    user,
	})
}

// Handler: Search user by name (query parameter)
func searchUser(c *gin.Context) {
	name := c.Query("name") // Ambil ?name=farhan

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query parameter 'name' is required",
		})
		return
	}

	// Dummy search
	if name == "farhan" || name == "Farhan" {
		user := User{ID: 1, Name: "Farhan", Email: "farhan@example.com"}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User found",
			"data":    user,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": "User not found with name: " + name,
	})
}