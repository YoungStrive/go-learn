package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int64
	id     int64
}

func main() {
	book1 := Book{Title: "GoLang", Author: "yanglin", Pages: 300, id: 1}
	fmt.Println(book1)

}
