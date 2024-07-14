package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Define structs to represent the JSON data
type Product struct {
	Peds           string      `json:"pedts"`
	IssuedTime     time.Time   `json:"issuedTime"`
	Wfo            string      `json:"wfo"`
	TimeZone       string      `json:"timeZone"`
	PrimaryName    string      `json:"primaryName"`
	PrimaryUnits   string      `json:"primaryUnits"`
	SecondaryName  string      `json:"secondaryName"`
	SecondaryUnits string      `json:"secondaryUnits"`
	Data           []DataPoint `json:"data"`
}

type DataPoint struct {
	ValidTime     time.Time `json:"validTime"`
	GeneratedTime time.Time `json:"generatedTime"`
	Primary       float32   `json:"primary"`
	Secondary     float32   `json:"secondary"`
}

func convertTimeToLocal(t time.Time) string {

	// Get the local location
	loc, err := time.LoadLocation("Local") // Use the system's local time zone
	if err != nil {
		return t.String()
	}

	// Convert the time to local time
	localTime := t.In(loc)

	// Format the local time in the desired format
	return localTime.Format("01/02 15:04")
}

func getJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}

func getProductForSite(siteId string, product string) (*Product, error) {
	if jsonData, err := getJSON(fmt.Sprintf("https://api.water.noaa.gov/nwps/v1/gauges/%s/stageflow/%s", siteId, product)); err != nil {
		fmt.Printf("Failed to get JSON: %v\n", err)
		return nil, err
	} else {
		// Unmarshal the JSON data into a Site struct
		var product Product
		err := json.Unmarshal(jsonData, &product)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil, err
		}
		return &product, nil
	}
}

func ObservedToString(siteId string) string {
	output := ""
	if observations, err := getProductForSite(siteId, "observed"); err != nil {
		output = fmt.Sprintf("Failed to get Site info: %v\n", err)
	} else {
		output = fmt.Sprintf("%s\nObservations issued at %s\n", output, convertTimeToLocal(observations.IssuedTime))
		count := 0
		for i := len(observations.Data) - 1; i >= 0; i-- {
			datum := observations.Data[i]
			output = fmt.Sprintf("%s%s %s:%2.2f %s:%2.2f\n", output,
				convertTimeToLocal(datum.ValidTime),
				observations.PrimaryUnits, datum.Primary,
				observations.SecondaryUnits, datum.Secondary)
			count++
			if count > 10 {
				break
			}
		}

	}
	return output
}

func ForecastToString(siteId string) string {
	output := ""
	if forecast, err := getProductForSite(siteId, "forecast"); err != nil {
		output = fmt.Sprintf("Failed to get Site info: %v\n", err)
	} else {
		output = fmt.Sprintf("%s\nForecast issued at %s\n", output, convertTimeToLocal(forecast.IssuedTime))
		for _, datum := range forecast.Data {
			output = fmt.Sprintf("%s%s %s:%2.2f %s:%2.2f\n", output,
				convertTimeToLocal(datum.ValidTime),
				forecast.PrimaryUnits, datum.Primary,
				forecast.SecondaryUnits, datum.Secondary)
		}

	}
	return output
}

// func main() {
// 	fmt.Print(ObservedToString("brkm2"))
// 	fmt.Print(ForecastToString("brkm2"))
// }
