package temperature

type fahrenheitHandler struct {
	Base int
}

// GetBase will return the word Celcius.
func (fh fahrenheitHandler) GetBase() string {
	return scaleToText[fh.Base]
}

// GetSymbol will return the Celcius symbol.
func (fh fahrenheitHandler) GetSymbol() string {
	return scaleToSymbol[fh.Base]
}

func (fh fahrenheitHandler) Convert() Handler {
	return fh
}

// ToCelcius returns the Celcius value for
// the provided temperature.
func (fh fahrenheitHandler) ToCelcius(temp float64) float64 {
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
	return fh.ToCelcius(temp) + KelvinBase
}

// NewFahrenheitConverter implements converter interface
// for Fahrenheit based temperatures.
func NewFahrenheitConverter() Converter {
	return &fahrenheitHandler{
		Base: Fahrenheit,
	}
}
