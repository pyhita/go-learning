package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// register routers
	// The biggest strength of the gorilla/mux Router is the ability to extract segments from the request URL
	r.HandleFunc("/books/{bookName}/pages/{pageNum}", func(w http.ResponseWriter, r *http.Request) {
		requests := mux.Vars(r)

		bookName := requests["bookName"]
		pageNum := requests["pageNum"]

		fmt.Fprintf(w, "The book is %s, page number is %s \n", bookName, pageNum)
	})

	http.ListenAndServe(":8888", r)
}
