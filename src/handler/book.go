package handler

import (
	"WebServiceGolang/src/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "David Pardamean Simatupang",
		"bio":  "Backend Developer and Software Engineer",
	})
}
func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":     "Web Service API GOLANG - David P Simatupang",
		"deskripsi": "Terimakasih",
	})
}

func (h *bookHandler) DetailProduk(c *gin.Context) {
	id := c.Param("id")
	judul := c.Param("judul")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"judul": judul,
	})
}

func (h *bookHandler) DetailQuery(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) PostBookhandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	//Ceking Error dalam Array
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {

			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return

	}
	//Jika tidak Terjadi Error Masuk Ke Database
	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"ID":          bookRequest.ID,
	// 	"title":       bookRequest.Title,
	// 	"price":       bookRequest.Price,
	// 	"description": bookRequest.Description,
	// 	"rating":      bookRequest.Rating,
	// 	"discount":    bookRequest.Discount,
	// })
}
