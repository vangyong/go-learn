package basic

import "fmt"

//结构体
type Book struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

//结构体作为函数参数传递
func PrintStruct(book Book) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book bookId : %d\n", book.BookId)
}
