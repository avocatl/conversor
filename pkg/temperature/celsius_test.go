package temperature

import (
	"testing"
)

var celsius = NewCelsiusConverter()

func TestGetBase_Celsius(t *testing.T) {
	if celsius.GetBase() != "Celsius" {
		t.Fail()
	}
}

func TestGetSymbol_Celsius(t *testing.T) {
	if celsius.GetSymbol() != "C" {
		t.Fail()
	}
}

func TestToCelsius_Celsius(t *testing.T) {
	cases := []struct {
		desc string
		want float64
	}{
		{
			"0 celsius to celsius",
			0,
		},
		{
			"1000 celsius to celsius",
			1000,
		},
		{
			"100 celsius to celsius",
			100,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celsius.Convert().ToCelsius(c.want)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToFahrenheit_Celsius(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 celsius to celsius",
			0,
			32,
		},
		{
			"100 celsius to celsius",
			100,
			212,
		},
		{
			"100 celsius to celsius",
			50,
			122,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celsius.Convert().ToFahrenheit(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToKelvin_Celsius(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 celsius to celsius",
			0,
			273.15,
		},
		{
			"100 celsius to celsius",
			100,
			373.15,
		},
		{
			"100 celsius to celsius",
			80,
			353.15,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celsius.Convert().ToKelvin(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}
