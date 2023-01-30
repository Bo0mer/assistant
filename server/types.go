package server

// Weather describes the weather for a specific place.
type Weather struct {
	// Description of the weather.
	Description string `json:"description"`
	// TempC holds the current temperature in Celsius.
	TempC int `json:"temp_c"`
}
