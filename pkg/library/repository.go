package library

import (
	"fmt"
	"log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Repository persistence layer for the Library
type Repository interface {
	Save(book Book) (int64, error)
	Delete(id int64) error
	Find(id int64) (*Book, error)
	Update(id int64, book Book) (int64, error)
	All() ([]Book, error)
}

// MysqlRepository Mysql Repository
type MysqlRepository struct {
	DB *sqlx.DB
}

// NewMysqlRepository returns a MysqlRepository
func NewMysqlRepository(conn string) *MysqlRepository {
	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	return &MysqlRepository{
		DB: db,
	}
}

// All get all books
func (r *MysqlRepository) All() ([]Book, error) {
	var books []Book
	err := r.DB.Select(&books, "SELECT title, author, year, isbn, description FROM library")
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Save save book
func (r *MysqlRepository) Save(book Book) (int64, error) {
	res := r.DB.MustExec("INSERT INTO library (title, isbn, year, author, description) VALUES (?,?,?,?,?)", book.Title, book.ISBN, book.Year, book.Author, book.Description)
	return res.LastInsertId()
}

// Delete book
func (r *MysqlRepository) Delete(id int64) error {
	res := r.DB.MustExec("DELETE FROM library WHERE id = ?", id)
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Book %d does not exists", id)
	}
	return nil
}

// Find book by id
func (r *MysqlRepository) Find(id int64) (*Book, error) {
	var book Book
	err := r.DB.Get(&book, "SELECT title, author, year, isbn, description FROM library WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Update book by id
func (r *MysqlRepository) Update(id int64, book Book) (int64, error) {
	res := r.DB.MustExec("UPDATE library SET title = ?, author = ?, year = ?, description = ?, isbn = ? WHERE id = ?", book.Title, book.Author, book.Year, book.Description, book.ISBN, id)
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return 0, fmt.Errorf("Book %d does not exists", id)
	}
	return id, nil
}
