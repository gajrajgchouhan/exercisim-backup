// Package weather returns the formatted string consisting of a location and its current weather conditions.
package weather

// CurrentCondition stores current weather conditions.
var CurrentCondition string
// CurrentLocation stores current location.
var CurrentLocation string

// Forecast returns the formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
