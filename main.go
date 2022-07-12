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
		// marshal request body into struct

		book := &Book{}

		err := c.ShouldBind(book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// use datastore to store struct
		err = datastore.CreateBook(c, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data": book,
		})
	})

	// Get Books
	server.GET("/books", func(c *gin.Context) {
		// marshal request body into struct
		books, err := datastore.FindBooks(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// use datastore to store struct

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":   &books,
		})
	})

	// Get Book
	server.GET("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		// https://pkg.go.dev/github.com/gin-gonic/gin

		// use datastore to store struct
		res, err := datastore.GetBook(c, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    res,
		})
	})

	// Update Book
	server.PATCH("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		err := c.ShouldBind(book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		id := c.Param("id") 
		// use datastore to store struct
		err = datastore.UpdateBook(c, id, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data": book,
		})
	})

	// Delete Book
	server.DELETE("/books/:id", func(c *gin.Context) {
		// marshal request body into struct

		book := &Book{}
		// use datastore to store struct
		err := datastore.DeleteBook(c, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"data":    book,
		})
	})
	log.Println("finished setting up")

	go func() {
		server.Run(fmt.Sprintf("%s", os.Getenv("PORT")))
	}()

	log.Println("server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("server closing")

}
