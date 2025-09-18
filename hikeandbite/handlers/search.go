// Package handlers contains handlers for REST API requests
package handlers

import (
	"encoding/json"
	"net/http"

	"hikeandbite/domain"
	"hikeandbite/services"
)

type RouteResponse struct {
	Name          string          `json:"name"`
	Distance      float32         `json:"distance"`
	ElevationGain int             `json:"elevationGain"`
	Cafes         []*CafeResponse `json:"cafes"`
}

type CafeResponse struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
}

func toRouteResponse(route *domain.Route) RouteResponse {
	var cafeResponses []*CafeResponse

	for _, cafe := range route.Cafes {
		r := &CafeResponse{
			Name:   cafe.Name,
			Rating: cafe.Rating,
		}
		cafeResponses = append(cafeResponses, r)
	}

	return RouteResponse{
		Name:          route.Name,
		Distance:      route.Distance,
		ElevationGain: route.ElevationGain,
		Cafes:         cafeResponses,
	}
}

type Handler struct {
	routeSearchService *services.RouteSearchService
}

func NewHandler(routeSearchService *services.RouteSearchService) Handler {
	return Handler{
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

	coord := domain.GeoCoordinates{Lat: lat, Lon: lon}
	routes := h.routeSearchService.SearchRoutes(coord)

	var responses []*RouteResponse

	for _, route := range routes {
		response := toRouteResponse(route)
		responses = append(responses, &response)
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(responses)
	if err != nil {
		http.Error(w, "Cannot serialize routes to JSON", http.StatusInternalServerError)
		return
	}
}
