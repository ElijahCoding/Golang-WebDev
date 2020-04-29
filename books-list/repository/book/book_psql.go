package bookRepository

import (
	"database/sql"
	"gofullstack/books-list/models"
	"log"
)

type BookRepository struct {}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	return books
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select * from books where id=$i", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}