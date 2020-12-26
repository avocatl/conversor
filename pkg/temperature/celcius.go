package temperature

type celciusHandler struct {
	Base int
}

// GetBase will return the word Celcius.
func (ch celciusHandler) GetBase() string {
	return scaleToText[ch.Base]
}

// GetSymbol will return the Celcius symbol.
func (ch celciusHandler) GetSymbol() string {
	return scaleToSymbol[ch.Base]
}

func (ch celciusHandler) Convert() Handler {
	return ch
}

// ToCelcius returns the Celcius value for
// the provided temperature.
func (ch celciusHandler) ToCelcius(temp float64) float64 {
	return temp
}

// ToFahrenheit returns the Fahrenheit value for
// the provided temperature.
func (ch celciusHandler) ToFahrenheit(temp float64) float64 {
	return (temp * fahrenheitMultiplier) + FahrenheitBase
}

// ToKelvin returns the Kelvin value for
// the provided temperature.
func (ch celciusHandler) ToKelvin(temp float64) float64 {
	return temp + KelvinBase
}

// NewCelciusConverter implements converter interface
// for Celcius based temperatures.
func NewCelciusConverter() Converter {
	return &celciusHandler{
		Base: Celcius,
	}
}
