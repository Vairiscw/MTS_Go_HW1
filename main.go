package main

import (
	"crypto/rand"
	"fmt"
)

type Book struct {
	name   string
	author string
	year   int
}

type idGenerator func(string) string

type Library struct {
	idGenerator idGenerator
	booksSlice  []Book
	booksMap    map[string]Book
}

func (l *Library) AddBook(book Book) string {
	id := l.idGenerator(book.name)
	l.booksMap[id] = book
	l.booksSlice = append(l.booksSlice, book)
	return id
}

func (l *Library) GetBookById(id string) Book {
	book, ok := l.booksMap[id]
	if ok {
		return book
	} else {
		return Book{}
	}
}

func (l *Library) GetBookByName(name string) Book {
	for _, book := range l.booksSlice {
		if book.name == name {
			return book
		}
	}
	return Book{}
}

func idGeneratorOne(key string) string {
	salt := rand.Text()
	return key + salt
}

func idGeneratorTwo(key string) string {
	return "id_" + key
}

func tests() {
	var bookSlice []Book
	bookSlice = append(bookSlice, Book{name: "Harry Potter", author: "J. K. Rowling", year: 1997})
	bookSlice = append(bookSlice, Book{name: "Курс теоретической физики", author: "Л. Д. Ландау", year: 1960})
	bookSlice = append(bookSlice, Book{name: "Grokking Algorithms", author: "Aditya Bhargava", year: 2016})

	library := Library{
		booksSlice:  make([]Book, 0),
		idGenerator: idGeneratorOne,
		booksMap:    make(map[string]Book),
	}

	var ids []string
	for _, book := range bookSlice {
		ids = append(ids, library.AddBook(book))
	}

	// Тестики
	if library.GetBookById(ids[0]).name != "Harry Potter" {
		fmt.Println("TEST1: ERROR")
	} else {
		fmt.Println("TEST1: OK")
	}
	if library.GetBookById(ids[2]).name != "Grokking Algorithms" {
		fmt.Println("TEST2: ERROR")
	} else {
		fmt.Println("TEST2: OK")
	}
	if library.GetBookByName("Курс теоретической физики").name != "Курс теоретической физики" {
		fmt.Println("TEST3: ERROR")
	} else {
		fmt.Println("TEST3: OK")
	}
	if library.GetBookByName("Grokking mainstream").name != "" {
		fmt.Println("TEST4: ERROR")
	} else {
		fmt.Println("TEST4: OK")
	}

	library.idGenerator = idGeneratorTwo
	library.booksSlice = make([]Book, 0)
	library.booksMap = make(map[string]Book)

	ids = make([]string, 0)
	for _, book := range bookSlice {
		ids = append(ids, library.AddBook(book))
	}
	if library.GetBookById(ids[0]).name != "Harry Potter" {
		fmt.Println("TEST5: ERROR")
	} else {
		fmt.Println("TEST5: OK")
	}
	if library.GetBookById(ids[2]).name != "Grokking Algorithms" {
		fmt.Println("TEST6: ERROR")
	} else {
		fmt.Println("TEST6: OK")
	}
}

func main() {
	tests()
}
