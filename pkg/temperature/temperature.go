package temperature

// Supported scales.
const (
	Celcius = iota
	Fahrenheit
	Kelvin
)

var (
	scaleToText = map[int]string{
		Celcius:    "Celcius",
		Fahrenheit: "Fahrenheit",
		Kelvin:     "Kelvin",
	}

	scaleToSymbol = map[int]string{
		Celcius:    "C",
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
	ToCelcius(temp float64) float64
	ToFahrenheit(temp float64) float64
	ToKelvin(temp float64) float64
}

// FahrenheitBase is the reference value added to any
// celcius temperature.
const FahrenheitBase float64 = 32

var fahrenheitMultiplier float64 = 1.8

// KelvinBase is the reference value added to any
// celcius temperature.
const KelvinBase float64 = 273.15
