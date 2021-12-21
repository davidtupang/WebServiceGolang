package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/produk/:id/:judul", detailProduk)
	v1.GET("/query/", detailQuery)
	v1.POST("/postbook", postBookhandler)

	router.Run()

}
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "David Pardamean Simatupang",
		"bio":  "Backend Developer and Software Engineer",
	})
}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":     "Web Service API GOLANG - David P Simatupang",
		"deskripsi": "Terimakasih",
	})
}

func detailProduk(c *gin.Context) {
	id := c.Param("id")
	judul := c.Param("judul")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"judul": judul,
	})
}

func detailQuery(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title" binding:"required"`
}

func postBookhandler(c *gin.Context) {
	var bookInput BookInput
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
