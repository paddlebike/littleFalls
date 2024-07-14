package main

import (
	// Shortening the import reference name seems to make it a bit easier
	"fmt"
	owm "github.com/briandowns/openweathermap"
)

var apiKey = "58f369e1efbff0ef7c1d8dce59ef4be2"

func oneCallByGeoPoint(lat float64, lon float64) string {
	// Possibility to exclude information. For example exclude daily information []string{ExcludeDaily}
	w, err := owm.NewOneCall("F", "EN", apiKey, []string{})
	if err != nil {
		return fmt.Sprintf("Error getting weather: %v", err)
	}
	err = w.OneCallByCoordinates(
		&owm.Coordinates{
			Longitude: lon,
			Latitude:  lat,
		},
	)
	if err != nil {
		return fmt.Sprintf("Error getting weather: %v", err)
	}
	c := w.Current
	output := fmt.Sprintf("Temp: %2.2f  FeelsLike: %2.2f Dew Point %2.2f\n", c.Temp, c.FeelsLike, c.DewPoint)
	output += fmt.Sprintf("Pressure: %d Humidity: %d Clouds: %d\n", c.Pressure, c.Humidity, c.Clouds)
	output += fmt.Sprintf("Wind Speed: %2.2f Gust %2.2f Direction %3.2f\n", c.WindSpeed, c.WindGust, c.WindDeg)

	if len(w.Alerts) > 0 {
		output += "\n Alerts"
		for _, v := range w.Alerts {
			output = fmt.Sprintf("%s\n%s\n", output, v.Event)
		}
	}

	return output
}
