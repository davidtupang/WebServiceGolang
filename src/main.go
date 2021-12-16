package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "David Pardamean Simatupang",
			"bio":  "Backend Developer and Software Engineer",
		})
	})
	router.Run()
}
