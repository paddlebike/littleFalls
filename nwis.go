package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type NWISGauges struct {
	SiteId string
	Time   time.Time
	Stage  float32
	Flow   uint32
	Temp   float32
}

func (n NWISGauges) toString() string {
	output := fmt.Sprintf("%s %2.2fc %dcfs %2.2ft",
		convertTimeToLocal(n.Time),
		n.Temp, n.Flow, n.Stage)
	return output
}

type NWISResponse struct {
	DeclaredType    string `json:"declaredType"`
	GlobalScope     bool   `json:"globalScope"`
	Name            string `json:"name"`
	Nil             bool   `json:"nil"`
	Scope           string `json:"scope"`
	TypeSubstituted bool   `json:"typeSubstituted"`
	NWISValue       struct {
		QueryInfo struct {
			Criteria struct {
				LocationParam string        `json:"locationParam"`
				Parameter     []interface{} `json:"parameter"`
				VariableParam string        `json:"variableParam"`
			} `json:"criteria"`
			Note []struct {
				Title string `json:"title"`
				Value string `json:"value"`
			} `json:"note"`
			QueryURL string `json:"queryURL"`
		} `json:"queryInfo"`
		Site []struct {
			Name       string `json:"name"`
			SourceInfo struct {
				GeoLocation struct {
					GeogLocation struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Srs       string  `json:"srs"`
					} `json:"geogLocation"`
					LocalSiteXY []interface{} `json:"localSiteXY"`
				} `json:"geoLocation"`
				Note     []interface{} `json:"note"`
				SiteCode []struct {
					AgencyCode string `json:"agencyCode"`
					Network    string `json:"network"`
					Value      string `json:"value"`
				} `json:"siteCode"`
				SiteName     string `json:"siteName"`
				SiteProperty []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"siteProperty"`
				SiteType     []interface{} `json:"siteType"`
				TimeZoneInfo struct {
					DaylightSavingsTimeZone struct {
						ZoneAbbreviation string `json:"zoneAbbreviation"`
						ZoneOffset       string `json:"zoneOffset"`
					} `json:"daylightSavingsTimeZone"`
					DefaultTimeZone struct {
						ZoneAbbreviation string `json:"zoneAbbreviation"`
						ZoneOffset       string `json:"zoneOffset"`
					} `json:"defaultTimeZone"`
					SiteUsesDaylightSavingsTime bool `json:"siteUsesDaylightSavingsTime"`
				} `json:"timeZoneInfo"`
			} `json:"sourceInfo"`
			TimeSeriesGauges []struct {
				CensorCode []interface{} `json:"censorCode"`
				Method     []struct {
					MethodDescription string `json:"methodDescription"`
					MethodID          int    `json:"methodID"`
				} `json:"method"`
				Offset    []interface{} `json:"offset"`
				Qualifier []struct {
					Network              string `json:"network"`
					QualifierCode        string `json:"qualifierCode"`
					QualifierDescription string `json:"qualifierDescription"`
					QualifierID          int    `json:"qualifierID"`
					Vocabulary           string `json:"vocabulary"`
				} `json:"qualifier"`
				QualityControlLevel []interface{} `json:"qualityControlLevel"`
				Sample              []interface{} `json:"sample"`
				Source              []interface{} `json:"source"`
				Value               []struct {
					DateTime   string   `json:"dateTime"`
					Qualifiers []string `json:"qualifiers"`
					Value      string   `json:"value"`
				} `json:"value"`
			} `json:"values"`
			Gauge struct {
				NoDataValue float64       `json:"noDataValue"`
				Note        []interface{} `json:"note"`
				Oid         string        `json:"oid"`
				Options     struct {
					Option []struct {
						Name       string `json:"name"`
						OptionCode string `json:"optionCode"`
					} `json:"option"`
				} `json:"options"`
				Unit struct {
					UnitCode string `json:"unitCode"`
				} `json:"unit"`
				Type string `json:"valueType"`
				Code []struct {
					Default    bool   `json:"default"`
					Network    string `json:"network"`
					Value      string `json:"value"`
					VariableID int    `json:"variableID"`
					Vocabulary string `json:"vocabulary"`
				} `json:"variableCode"`
				Description string        `json:"variableDescription"`
				Name        string        `json:"variableName"`
				Property    []interface{} `json:"variableProperty"`
			} `json:"variable"`
		} `json:"timeSeries"`
	} `json:"value"`
}

func (ts NWISResponse) GaugeTime(index int) time.Time {
	gauges := ts.NWISValue.Site[index].TimeSeriesGauges
	lastReading := gauges[len(gauges)-1].Value[0].DateTime
	gt, err := time.Parse(time.RFC3339, lastReading)
	if err != nil {
		return time.Now()
	}
	return gt
}

func (ts NWISResponse) getNWISGauges(index int) (*NWISGauges, error) {
	ng := NWISGauges{
		SiteId: ts.NWISValue.Site[0].SourceInfo.SiteCode[0].Value,
		Time:   ts.GaugeTime(index),
	}
	for _, v := range ts.NWISValue.Site {
		valueString := v.TimeSeriesGauges[index].Value[0].Value
		switch v.Gauge.Oid {
		case "45807042":
			c, err := strconv.ParseFloat(valueString, 32)
			if err == nil {
				ng.Temp = float32(c)
			}
		case "45807197":
			c, err := strconv.ParseUint(valueString, 10, 32)
			if err == nil {
				ng.Flow = uint32(c)
			}
		case "45807202":

			c, err := strconv.ParseFloat(valueString, 32)
			if err == nil {
				ng.Stage = float32(c)
			}
		}
	}
	return &ng, nil
}

func (ts NWISResponse) getNWISGaugeCount() int {
	return len(ts.NWISValue.Site)
}

func (ts NWISResponse) GaugeString(index int) string {
	gauges := ts.NWISValue.Site[index].TimeSeriesGauges
	lastReading := gauges[len(gauges)-1].Value[0].Value
	// Get the name
	name := "unknown"
	value := "Unkown"
	switch ts.NWISValue.Site[index].Gauge.Oid {
	case "45807042":
		name = "Temperature"
		c, err := strconv.ParseFloat(lastReading, 32)
		if err == nil {
			f := (c * 1.8) + 32
			value = fmt.Sprintf("%2.2fC : %3.2fF", c, f)
		}
	case "45807197":
		name = "Height FT"
		value = lastReading
	case "45807202":
		name = "Flow CFS"
		value = lastReading
	}

	return fmt.Sprintf("%-15s: %s", name, value)
}

func (ts NWISResponse) SiteName() string {

	return ts.NWISValue.Site[0].SourceInfo.SiteName
}

func (ts NWISResponse) ToString() string {

	siteName := ts.SiteName()

	readingTime := ts.GaugeTime(0)
	v0 := ts.GaugeString(0)
	v1 := ts.GaugeString(1)
	v2 := ts.GaugeString(2)

	output := fmt.Sprintf("%s\nConditions AT: %s\n%s\n%s\n%s", siteName, readingTime, v0, v1, v2)
	for i := 0; i < ts.getNWISGaugeCount(); i++ {
		g, err := ts.getNWISGauges(0)
		if g != nil && err == nil {
			output = output + "\n" + g.toString()
		}
	}
	return output
}

func getNWISSite(siteId string) (*NWISResponse, error) {
	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&parameterCd=00065,00060,00010&sites=%s", siteId)
	if jsonData, err := getJSON(url); err != nil {
		fmt.Printf("Failed to get JSON: %v\n", err)
		return nil, err
	} else {
		// Unmarshal the JSON data into a Site struct
		var ts NWISResponse
		err := json.Unmarshal(jsonData, &ts)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil, err
		}
		return &ts, nil
	}
}
