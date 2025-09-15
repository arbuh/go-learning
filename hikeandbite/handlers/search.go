// Package handlers contains handlers for REST API requests
package handlers

import (
	"fmt"
	"net/http"

	"hikeandbite/domain"
	"hikeandbite/services"
)

type RouteResponse struct {
	Name          string
	Distance      float32
	ElevationGain int
	Cafes         []*CafeResponse
}

type CafeResponse struct {
	Name   string
	Rating float32
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

	routes := h.routeSearchService.SearchRoutes(lat, lon)

	var responses []*RouteResponse

	for _, route := range routes {
		response := toRouteResponse(route)
		responses = append(responses, &response)
	}

	fmt.Fprintf(w, "Coordinates received: lat=%s, lon=%s", lat, lon)
}
