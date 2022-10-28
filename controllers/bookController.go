package controllers

// Imports

import (
	"github.com/cap811/go-rest-api/initializers"
	"github.com/cap811/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET Request
// Show one book

func FindBook(c *gin.Context) {
	var book []models.Book
	id := c.Param("id")
	initializers.DB.First(&book, id)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET Request
// Show all books

func ShowAllBooks(c *gin.Context) {
	var books []models.Book
	initializers.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST Request
// Create books

func CreateBooks(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author, Year: input.Year}
	initializers.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST Request
// Update book

func UpdateBooks(c *gin.Context) {
	var input models.UpdateBookInput
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author, Year: input.Year}
	initializers.DB.Where(id).Updates(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
