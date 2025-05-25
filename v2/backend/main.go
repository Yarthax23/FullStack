package main

import (
	"net/http"

	"github.com/Yarthax23/FullStack/tree/main/v2/backend/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)

	http.ListenAndServe("8090", nil)
}
