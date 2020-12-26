package temperature

// Supported scales.
const (
	Celsius = iota
	Fahrenheit
	Kelvin
)

var (
	scaleToText = map[int]string{
		Celsius:    "Celsius",
		Fahrenheit: "Fahrenheit",
		Kelvin:     "Kelvin",
	}

	scaleToSymbol = map[int]string{
		Celsius:    "C",
		Fahrenheit: "F",
		Kelvin:     "K",
	}
)

// Converter interface describes a temperature
// conversion handler.
type Converter interface {
	GetBase() string
	GetSymbol() string
	Convert() Handler
}

// Handler interface describes a temperature
// conversion handler that simplifies the
// converter implementation.
type Handler interface {
	ToCelsius(temp float64) float64
	ToFahrenheit(temp float64) float64
	ToKelvin(temp float64) float64
}

// FahrenheitBase is the reference value added to any
// celsius temperature.
const FahrenheitBase float64 = 32

var fahrenheitMultiplier float64 = 1.8

// KelvinBase is the reference value added to any
// celsius temperature.
const KelvinBase float64 = 273.15
