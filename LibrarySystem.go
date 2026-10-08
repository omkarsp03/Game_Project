package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Issued bool
}

type Member struct {
	ID   int
	Name string
}

type LibraryItem interface {
	Display()
	Issue()
	Return()
}

func (b *Book) Display() {
	fmt.Println("ID:", b.ID)
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
}

func (b *Book) Issue() {
	if b.Issued {
		fmt.Println("Book is already issued")
		return
	}

	b.Issued = true
	fmt.Println("Book issued successfully")
}

func (b *Book) Return() {
	b.Issued = false
	fmt.Println("Book returned successfully")
}

func processItem(item LibraryItem) {
	item.Display()
	item.Issue()
}

func main() {
	book := Book{
		ID:     1,
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Issued: false,
	}

	processItem(&book)

	book.Return()
}
