package main

import (
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.ListenAndServe(":8080", nil)
}
