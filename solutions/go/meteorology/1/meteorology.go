package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (sp SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sp]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (sp Speed) String() string {
	return fmt.Sprintf("%v %v", sp.magnitude, sp.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (mData MeteorologyData) String() string {
	return fmt.Sprintf("%s: %v, Wind %v at %v, %v%% Humidity", mData.location, mData.temperature,
		mData.windDirection, mData.windSpeed, mData.humidity)
}
