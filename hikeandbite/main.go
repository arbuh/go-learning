package main

import (
	"fmt"
	"net/http"

	"hikeandbite/handlers"
	"hikeandbite/services"
)

func main() {
	routeSearchService := services.NewRouteSearchService()
	handlers := handlers.NewHandler(routeSearchService)

	http.HandleFunc("/search", handlers.SearchByCoordinates)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error", err)
	}
}
