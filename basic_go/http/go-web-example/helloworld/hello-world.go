package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, you've requested %s \n", request.URL)
	})

	http.ListenAndServe(":8888", nil)
}
