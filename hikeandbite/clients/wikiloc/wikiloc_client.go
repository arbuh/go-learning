// Package wikiloc contains clients for Wikiloc.com
package wikiloc

import (
	"hikeandbite/domain"
	"net/http"
	"time"
)

type WikilocClient struct {
	httpClient *http.Client
}

func NewWikilocClient() *WikilocClient {
	return &WikilocClient{httpClient: &http.Client{Timeout: 30 * time.Second}}
}

// curl 'https://www.wikiloc.com/wikiloc/find.do?event=map&to=24&sw=52.046261923665206%2C5.702548027038575&ne=52.047053736959754%2C5.750269889831544' \
//  -H 'accept: */*' \
//  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36'
func (clint WikilocClient) SearchRoutes(coord domain.GeoCoordinates) []domain.Route {
	return nil
}
