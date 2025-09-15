// Package domain contains shared business logic entities
package domain

type Route struct {
	Name          string
	Link          string
	Distance      float32
	ElevationGain int
	Cafes         []*Cafe
}
