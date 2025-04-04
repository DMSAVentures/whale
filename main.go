package main

import (
	"league/internal/api"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrixoperations.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", api.EchoHandler)
	http.HandleFunc("/invert", api.InvertHandler)
	http.HandleFunc("/sum", api.SumHandler)
	http.HandleFunc("/multiply", api.MultiplyHandler)
	http.HandleFunc("/flatten", api.FlattenHandler)
	http.ListenAndServe(":8080", nil)
}
