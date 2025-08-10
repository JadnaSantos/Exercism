// Package weather provides methods for weather reporting.
package weather

// CurrentCondition is the local wheather condition.
var CurrentCondition string

// CurrentLocation is the local wheather location.
var CurrentLocation string

// Forecast return a string value to the weather currentLocation and currentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
