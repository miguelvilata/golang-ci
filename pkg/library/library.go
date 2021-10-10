package library

// Library struct
type Library struct {
	Repository Repository
}

// NewLibrary creates a new Library
func NewLibrary(repository Repository) *Library {
	return &Library{
		Repository: repository,
	}
}

// All get all books
func (library *Library) All() (books []Book, err error) {
	return library.Repository.All()
}

// Add adds a new book to the Library
func (library *Library) Add(book Book) (id int64, err error) {
	return library.Repository.Save(book)
}

// Delete removes a book from the Library
func (library *Library) Delete(id int64) error {
	return library.Repository.Delete(id)
}

// Find find a book by id from the Library
func (library *Library) Find(id int64) (book *Book, err error) {
	return library.Repository.Find(id)
}

// Modify a book
func (library *Library) Modify(id int64, book Book) (int64, error) {
	return library.Repository.Update(id, book)
}
