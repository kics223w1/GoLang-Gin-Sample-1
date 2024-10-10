package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quanity int `json:"quanity"`
}

var books = []book{
	{ID: "1", Title:"In search", Author: "Huyhuy", Quanity: 10},
	{ID: "2", Title:"In search", Author: "Huyhuy", Quanity: 10},
	{ID: "3", Title:"In search", Author: "Huyhuy", Quanity: 12 },
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}


func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id , ok := c.GetQuery("id");

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" :"Missing query."});
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" :"Book not found."});
		return
	}

	if book.Quanity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" :"Out of book."});
		return
	}

	book.Quanity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id , ok := c.GetQuery("id");

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" :"Missing query."});
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" :"Book not found."});
		return
	}

	book.Quanity += 1
	c.IndentedJSON(http.StatusOK, book)
}


func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout" , checkoutBook)
	router.PATCH("/checkin" , returnBook)
	router.Run("localhost:8080")
}