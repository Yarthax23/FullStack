package main

import (
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)

	http.ListenAndServe("8090", nil)
}
