package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(id int, book BookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	//return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
	//return s.repository.FindAll()
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}
	newbook, err := s.repository.Create(book)
	return newbook, err

}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindById(id)

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating
	book.Discount = bookRequest.Discount

	newbook, err := s.repository.Update(book)
	return newbook, err

}

func (s *service) Delete(id int) (Book, error) {
	book, _ := s.repository.FindById(id)

	newbook, err := s.repository.Delete(book)
	return newbook, err

}
