package main

import "fmt"

type Book struct {
	title string
	author string
	price float32
}

func main() {
	b1 := Book {
		title: "golang",
		author: "zhangsan",
		price: 22.9,
	}

	fmt.Printf("book: %v \n", b1)
	fmt.Println("book title: ", b1.title, "book author: ", b1.author, "book price: ", b1.price)
}