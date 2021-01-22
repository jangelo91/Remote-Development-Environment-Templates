package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		io.WriteString(w, "Hello")
		//fmt.Fprintf(w, "Welcome to my website!")
	})

	http.ListenAndServe(":9000", nil)
}
