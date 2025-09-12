package main

import (
	"fmt"
	"net/http"

	"hikeandbite/handlers"
)

func main() {
	http.HandleFunc("/search", handlers.SearchByCoordinates)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error", err)
	}
}
