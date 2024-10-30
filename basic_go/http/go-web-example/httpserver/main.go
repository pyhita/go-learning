package main

import (
	"fmt"
	"net/http"
)

func main() {

	// test?username=xxxx&password=xxx
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// parse params from request
		query := r.URL.Query()
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		fmt.Println("query: ", query, "username: ", username, "password: ", password)
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// register file server
	fs := http.FileServer(http.Dir("/Users/pyhita/Codes/go-codes/go-learning/basic_go/http/go-web-example/httpserver/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8888", nil)
}
