// Package handlers contains handlers for REST API requests
package handlers

import (
	"fmt"
	"net/http"
)

func SearchByCoordinates(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "Coordinates received: lat=%s, lon=%s", lat, lon)
}
