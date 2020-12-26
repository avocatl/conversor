package temperature

type kelvinHandler struct {
	Base int
}

// GetBase will return the word Celcius.
func (kh kelvinHandler) GetBase() string {
	return scaleToText[kh.Base]
}

// GetSymbol will return the Celcius symbol.
func (kh kelvinHandler) GetSymbol() string {
	return scaleToSymbol[kh.Base]
}

func (kh kelvinHandler) Convert() Handler {
	return kh
}

// ToCelcius returns the Celcius value for
// the provided temperature.
func (kh kelvinHandler) ToCelcius(temp float64) float64 {
	return temp - KelvinBase
}

// ToFahrenheit returns the Fahrenheit value for
// the provided temperature.
func (kh kelvinHandler) ToFahrenheit(temp float64) float64 {
	return (kh.ToCelcius(temp) * fahrenheitMultiplier) + FahrenheitBase
}

// ToKelvin returns the Kelvin value for
// the provided temperature.
func (kh kelvinHandler) ToKelvin(temp float64) float64 {
	return temp
}

// NewKelvinConverter implements converter interface
// for Kelvin based temperatures.
func NewKelvinConverter() Converter {
	return &kelvinHandler{
		Base: Kelvin,
	}
}
