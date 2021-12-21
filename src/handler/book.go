package handler

import (
	"WebServiceGolang/src/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "David Pardamean Simatupang",
		"bio":  "Backend Developer and Software Engineer",
	})
}
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":     "Web Service API GOLANG - David P Simatupang",
		"deskripsi": "Terimakasih",
	})
}

func DetailProduk(c *gin.Context) {
	id := c.Param("id")
	judul := c.Param("judul")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"judul": judul,
	})
}

func DetailQuery(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBookhandler(c *gin.Context) {
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)

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

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
