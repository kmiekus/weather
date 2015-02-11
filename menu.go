package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1235", nil)
}

func handler(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Write([]byte("test"))
}
