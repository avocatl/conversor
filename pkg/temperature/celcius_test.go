package temperature

import (
	"testing"
)

var celcius = NewCelciusConverter()

func TestGetBase_Celcius(t *testing.T) {
	if celcius.GetBase() != "Celcius" {
		t.Fail()
	}
}

func TestGetSymbol_Celcius(t *testing.T) {
	if celcius.GetSymbol() != "C" {
		t.Fail()
	}
}

func TestToCelcius_Celcius(t *testing.T) {
	cases := []struct {
		desc string
		want float64
	}{
		{
			"0 celcius to celcius",
			0,
		},
		{
			"1000 celcius to celcius",
			1000,
		},
		{
			"100 celcius to celcius",
			100,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celcius.Convert().ToCelcius(c.want)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToFahrenheit_Celcius(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 celcius to celcius",
			0,
			32,
		},
		{
			"100 celcius to celcius",
			100,
			212,
		},
		{
			"100 celcius to celcius",
			50,
			122,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celcius.Convert().ToFahrenheit(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToKelvin_Celcius(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 celcius to celcius",
			0,
			273.15,
		},
		{
			"100 celcius to celcius",
			100,
			373.15,
		},
		{
			"100 celcius to celcius",
			80,
			353.15,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := celcius.Convert().ToKelvin(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}
