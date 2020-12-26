package temperature

import (
	"log"
	"testing"
)

var kelvin = NewKelvinConverter()

func TestGetBase_Kelvin(t *testing.T) {
	if kelvin.GetBase() != "Kelvin" {
		t.Fail()
	}
}

func TestGetSymbol_Kelvin(t *testing.T) {
	if kelvin.GetSymbol() != "K" {
		t.Fail()
	}
}

func TestToCelcius_Kelvin(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 kelvin to celcius",
			0,
			-273.15,
		},
		{
			"100 kelvin to celcius",
			100,
			-173.14999999999998,
		},
		{
			"32 kelvin to celcius",
			293.15,
			20,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := kelvin.Convert().ToCelcius(c.temp)

			log.Println(got)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToFahrenheit_Kelvin(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 kelvin to fahrenheit",
			0,
			-459.66999999999996,
		},
		{
			"100 kelvin to fahrenheit",
			100,
			-279.66999999999996,
		},
		{
			"300 kelvin to fahrenheit",
			300,
			80.33000000000004,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := kelvin.Convert().ToFahrenheit(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToKelvin_Kelvin(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 kelvin to kelvin",
			0,
			0,
		},
		{
			"100 kelvin to kelvin",
			100,
			100,
		},
		{
			"100 kelvin to kelvin",
			80,
			80,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := kelvin.Convert().ToKelvin(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}
