// Package weather provides a forecast.
package weather

var (
	// CurrentCondition represents the current weather.
	CurrentCondition string
	// CurrentLocation represntes the location.
	CurrentLocation string
)

// Forecast returns the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
