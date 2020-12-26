package temperature

type celsiusHandler struct {
	Base int
}

// GetBase will return the word Celsius.
func (ch celsiusHandler) GetBase() string {
	return scaleToText[ch.Base]
}

// GetSymbol will return the Celsius symbol.
func (ch celsiusHandler) GetSymbol() string {
	return scaleToSymbol[ch.Base]
}

func (ch celsiusHandler) Convert() Handler {
	return ch
}

// ToCelsius returns the Celsius value for
// the provided temperature.
func (ch celsiusHandler) ToCelsius(temp float64) float64 {
	return temp
}

// ToFahrenheit returns the Fahrenheit value for
// the provided temperature.
func (ch celsiusHandler) ToFahrenheit(temp float64) float64 {
	return (temp * fahrenheitMultiplier) + FahrenheitBase
}

// ToKelvin returns the Kelvin value for
// the provided temperature.
func (ch celsiusHandler) ToKelvin(temp float64) float64 {
	return temp + KelvinBase
}

// NewCelsiusConverter implements converter interface
// for Celsius based temperatures.
func NewCelsiusConverter() Converter {
	return &celsiusHandler{
		Base: Celsius,
	}
}
