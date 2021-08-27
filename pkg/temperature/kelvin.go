package temperature

type kelvinHandler struct {
	Base int
}

// GetBase will return the word Celsius.
func (kh kelvinHandler) GetBase() string {
	return scaleToText(kh.Base)
}

// GetSymbol will return the Celsius symbol.
func (kh kelvinHandler) GetSymbol() string {
	return scaleToSymbol(kh.Base)
}

func (kh kelvinHandler) Convert() Handler {
	return kh
}

// ToCelsius returns the Celsius value for
// the provided temperature.
func (kh kelvinHandler) ToCelsius(temp float64) float64 {
	return temp - KelvinBase
}

// ToFahrenheit returns the Fahrenheit value for
// the provided temperature.
func (kh kelvinHandler) ToFahrenheit(temp float64) float64 {
	return (kh.ToCelsius(temp) * fahrenheitMultiplier) + FahrenheitBase
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
