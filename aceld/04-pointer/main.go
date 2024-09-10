package main

import "fmt"

func swap(a *int, b *int) {
	fmt.Println("a = ", a, " b = ", b)
	temp := *a
	*a = *b
	*b = temp
}


func main() {	
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)
}