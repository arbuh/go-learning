// Package handlers contains handlers for REST API requests
package handlers

import (
	"fmt"
	"net/http"

	"hikeandbite/services"
)

type Handler struct {
	routeSearchService *services.RouteSearchService
}

func NewHandler(routeSearchService *services.RouteSearchService) Handler {
	return Handler {
		routeSearchService: routeSearchService,
	}
}

func (h *Handler) SearchByCoordinates(w http.ResponseWriter, r *http.Request) {
	lon := r.URL.Query().Get("lon")
	lat := r.URL.Query().Get("lat")

	if lon == "" {
		http.Error(w, "The query parameter 'lon' is mandatory", http.StatusBadRequest)
		return
	}

	if lat == "" {
		http.Error(w, "The query parameter 'lat' is mandatory", http.StatusBadRequest)
		return
	}

	h.routeSearchService.SearchRoutes(lat, lon)

	fmt.Fprintf(w, "Coordinates received: lat=%s, lon=%s", lat, lon)
}
