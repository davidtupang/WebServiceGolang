package main

import (
	"WebServiceGolang/src/book"
	"WebServiceGolang/src/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3307)/servicegolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	db.AutoMigrate(&book.Book{})
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/produk/:id/:judul", bookHandler.DetailProduk)
	v1.GET("/query/", bookHandler.DetailQuery)
	v1.POST("/postbook", bookHandler.PostBookhandler)

	router.Run(":99")

	//Service Testing
	// bookRepository := book.NewRepository(db)
	// bookService := book.NewService(bookRepository)
	// bookRequest := book.BookRequest{
	// 	Title:       "Halo",
	// 	Description: "Mantap",
	// 	Price:       75000,
	// 	Rating:      5,
	// 	Discount:    25,
	// }
	// bookService.Create(bookRequest)

	//REPOSITORY Testing

	//bookRepository := book.NewRepository(db)

	//FINDALL Repo
	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// }

	//FINDBYID Repo
	// id, err := bookRepository.FindById(2)
	// fmt.Println("Title : ", id.Title)

	//CREATE Repo
	// book := book.Book{
	// 	Title:       "Heroik",
	// 	Description: "Good Book",
	// 	Price:       95000,
	// 	Rating:      4,
	// 	Discount:    0,
	// }
	// bookRepository.Create(book)

	//=================================

	//CRUD Testing

	// ======
	// CREATE
	// ======
	// book := book.Book{}
	// book.Title = "Atomic habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Self Development Kebiasaan Baik"
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Creating Book")
	// 	fmt.Println("===================")
	// }

	// ======
	// READ
	// ======
	// var books []book.Book

	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	// 	fmt.Println("===================")
	// 	// 	fmt.Println("Error Finding Book")
	// 	// 	fmt.Println("===================")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("Book Objek %v :", b)
	// }

	//UPDATE
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).Find(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Finding Book")
	// 	fmt.Println("===================")
	// }
	// book.Title = "Man Tiger (Revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Updating Book")
	// 	fmt.Println("===================")
	// }

	//DELETE
	// var book book.Book
	// err = db.Debug().Where("id", 3).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Finding Book")
	// 	fmt.Println("===================")
	// }
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Deleting Book Book")
	// 	fmt.Println("===================")
	// }

	// Testing Routing
	// router := gin.Default()

	// v1 := router.Group("/v1")

	// v1.GET("/", handler.RootHandler)
	// v1.GET("/hello", handler.HelloHandler)
	// v1.GET("/produk/:id/:judul", handler.DetailProduk)
	// v1.GET("/query/", handler.DetailQuery)
	// v1.POST("/postbook", handler.PostBookhandler)

	// router.Run()

}
