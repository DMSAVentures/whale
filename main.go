package main

import (
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrixoperations.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.ListenAndServe(":8080", nil)
}
