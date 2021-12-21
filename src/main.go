package main

import (
	"WebServiceGolang/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/produk/:id/:judul", handler.DetailProduk)
	v1.GET("/query/", handler.DetailQuery)
	v1.POST("/postbook", handler.PostBookhandler)

	router.Run()

}
