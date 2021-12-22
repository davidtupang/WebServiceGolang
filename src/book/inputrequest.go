package book

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Discount    int    `json:"discount" binding:"required"`
}
