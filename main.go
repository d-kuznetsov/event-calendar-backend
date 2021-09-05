package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Some message")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
