package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	fmt.Println("mongo: ", mongoURI)

	ctx := context.Background()
	log.Println("create datastore")

	datastore := NewDatastore(ctx, mongoURI)
	_ = datastore

	log.Println("setup server datastore")
	server := gin.Default()

	// Create Book
	server.POST("/books", func(c *gin.Context) {
		// marshclal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get Books
	server.GET("/books", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get Book
	server.GET("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Update Book
	server.PATCH("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Delete Book
	server.DELETE("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	log.Println("finished setting up")

	go func() {
		server.Run("localhost:8080")
	}()

	log.Println("server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("server closing")

}
