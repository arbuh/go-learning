package main

import (
	"fmt"
	"math"
	"strconv"
)

type CheapestFlight struct {
	Price int
	Day   int
}

func searchFlights(startDay int, endDay int) map[int]int {
	mockFlights := make(map[int]int)

	price := 100

	for startDay <= endDay {
		mockFlights[startDay] = price

		price += 50
		startDay++
	}

	return mockFlights
}

func searchCheapestFlight(flightRange int, startDay int, endDay int) map[string]CheapestFlight {
	cheapestFlights := make(map[string]CheapestFlight)

	flights := searchFlights(startDay, endDay)

	for i := startDay; i <= endDay-flightRange+1; i++ {
		rangeStart := i
		rangeEnd := i + flightRange - 1

		minPrice := math.MaxInt
		minPriceDay := -1

		for j := rangeStart; j <= rangeEnd; j++ {
			if flights[j] < minPrice {
				minPrice = flights[j]
				minPriceDay = j
			}
		}

		key := strconv.Itoa(rangeStart) + "-" + strconv.Itoa(rangeEnd)
		cheapestFlights[key] = CheapestFlight{Price: minPrice, Day: minPriceDay}
	}

	return cheapestFlights
}

func main() {
	cheapestFlights := searchCheapestFlight(3, 1, 9)

	for key, flight := range cheapestFlights {
		fmt.Printf("%s: the minimal price is %d on the day %d\n", key, flight.Price, flight.Day)
	}
}
