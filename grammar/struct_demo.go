package grammar

import "fmt"

//结构体
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

//结构体作为函数参数传递
func printStruct(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book bookId : %d\n", book.bookId)
}
