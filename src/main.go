package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/produk/:id/:judul", detailProduk)
	router.GET("/query/", detailQuery)
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
