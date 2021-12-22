package handler

import (
	"WebServiceGolang/src/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// func (h *bookHandler) RootHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "David Pardamean Simatupang",
// 		"bio":  "Backend Developer and Software Engineer",
// 	})
// }
// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"title":     "Web Service API GOLANG - David P Simatupang",
// 		"deskripsi": "Terimakasih",
// 	})
// }

// func (h *bookHandler) DetailProduk(c *gin.Context) {
// 	id := c.Param("id")
// 	judul := c.Param("judul")
// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"judul": judul,
// 	})
// }

// func (h *bookHandler) DetailQuery(c *gin.Context) {
// 	id := c.Query("id")
// 	title := c.Query("title")
// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

func (h *bookHandler) GetBookAllhandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})

}

func (h *bookHandler) GetBookByIdhandler(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)
	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	bookResponse := convertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
func (h *bookHandler) UpdateBookByIdhandler(c *gin.Context) {

}
func (h *bookHandler) DeleteBookByIdhanler(c *gin.Context) {

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
	postResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": postResponse,
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

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
