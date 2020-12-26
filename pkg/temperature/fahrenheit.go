package temperature

type fahrenheitHandler struct {
	Base int
}

// GetBase will return the word Celsius.
func (fh fahrenheitHandler) GetBase() string {
	return scaleToText[fh.Base]
}

// GetSymbol will return the Celsius symbol.
func (fh fahrenheitHandler) GetSymbol() string {
	return scaleToSymbol[fh.Base]
}

func (fh fahrenheitHandler) Convert() Handler {
	return fh
}

// ToCelsius returns the Celsius value for
// the provided temperature.
func (fh fahrenheitHandler) ToCelsius(temp float64) float64 {
	return (temp - FahrenheitBase) * 5 / 9
}

// ToFahrenheit returns the Fahrenheit value for
// the provided temperature.
func (fh fahrenheitHandler) ToFahrenheit(temp float64) float64 {
	return temp
}

// ToKelvin returns the Kelvin value for
// the provided temperature.
func (fh fahrenheitHandler) ToKelvin(temp float64) float64 {
	return fh.ToCelsius(temp) + KelvinBase
}

// NewFahrenheitConverter implements converter interface
// for Fahrenheit based temperatures.
func NewFahrenheitConverter() Converter {
	return &fahrenheitHandler{
		Base: Fahrenheit,
	}
}
