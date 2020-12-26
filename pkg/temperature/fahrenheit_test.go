package temperature

import (
	"log"
	"testing"
)

var fahrenheit = NewFahrenheitConverter()

func TestGetBase_Fahrenheit(t *testing.T) {
	if fahrenheit.GetBase() != "Fahrenheit" {
		t.Fail()
	}
}

func TestGetSymbol_Fahrenheit(t *testing.T) {
	if fahrenheit.GetSymbol() != "F" {
		t.Fail()
	}
}

func TestToCelsius_Fahrenheit(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 fahrenheit to celsius",
			0,
			-17.77777777777778,
		},
		{
			"100 fahrenheit to celsius",
			100,
			37.77777777777778,
		},
		{
			"32 fahrenheit to celsius",
			32,
			0,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := fahrenheit.Convert().ToCelsius(c.temp)

			log.Println(got)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToFahrenheit_Fahrenheit(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 fahrenheit to fahrenheit",
			32,
			32,
		},
		{
			"100 fahrenheit to fahrenheit",
			212,
			212,
		},
		{
			"100 fahrenheit to fahrenheit",
			122,
			122,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := fahrenheit.Convert().ToFahrenheit(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}

func TestToKelvin_Fahrenheit(t *testing.T) {
	cases := []struct {
		desc string
		temp float64
		want float64
	}{
		{
			"0 fahrenheit to kelvin",
			0,
			255.3722222222222,
		},
		{
			"100 fahrenheit to kelvin",
			100,
			310.92777777777775,
		},
		{
			"100 fahrenheit to kelvin",
			80,
			299.81666666666666,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := fahrenheit.Convert().ToKelvin(c.temp)

			if got != c.want {
				t.Errorf("conversion failed, got %f, want %f", got, c.want)
			}
		})
	}
}
