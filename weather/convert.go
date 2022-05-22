package weather

// KelvinToCelcius converts temperature in Kelvin to Celcius.
func KelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}
