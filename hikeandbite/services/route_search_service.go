// Package services contains the main application logic that involves multiple domain entities and managing of side effects
package services

import "hikeandbite/domain"

func SearchRoutes(lat string, lon string) []domain.Route {
	cafe1 := domain.Cafe{
		Name:   "Restaurant Planken Wambuis",
		Link:   "https://maps.app.goo.gl/o4CQFnUoJaUxQHTR6",
		Rating: 4.4,
	}

	cafe2 := domain.Cafe{
		Name:   "Boerderij Mossel",
		Link:   "https://maps.app.goo.gl/JQc9yUrawHidntNx9",
		Rating: 4.5,
	}

	route1 := domain.Route{
		Name:          "Mossel, Ginkelse heide en Planken Wambuis",
		Link:          "https://www.wikiloc.com/hiking-trails/mossel-ginkelse-heide-en-planken-wambuis-190768873",
		Distance:      12.0,
		ElevationGain: 100,
		Cafes:         []domain.Cafe{cafe1, cafe2},
	}

	route2 := domain.Route{
		Name:          "Rondje lopen vanaf Boerderij Mossel",
		Link:          "https://www.wikiloc.com/hiking-trails/rondje-lopen-vanaf-boerderij-mossel-150526201",
		Distance:      10.0,
		ElevationGain: 80,
		Cafes:         []domain.Cafe{cafe1, cafe2},
	}

	return []domain.Route{route1, route2}
}
