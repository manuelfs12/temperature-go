package main

import "testing"

func TestConvertCelsiusToKelvin(t *testing.T) {
	got := convertCelsiusToKelvin(50)
	want := "323.15"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	got := convertCelsiusToFahrenheit(50)
	want := "122.00"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestConvertKelvinToCelsius(t *testing.T) {
	got := convertKelvinToCelsius(50)
	want := "-223.15"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestConvertKelvinToFahrenheit(t *testing.T) {
	got := convertKelvinToFahrenheit(50)
	want := "-369.67"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestConvertFahrenheitToKelvin(t *testing.T) {
	got := convertFahrenheitToKelvin(50)
	want := "283.15"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestConvertFahrenheitToCelsius(t *testing.T) {
	got := convertFahrenheitToCelsius(50)
	want := "10.00"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
