package main

import (
	"encoding/json"
	"regexp"
	"testing"
)

var nwisJSON = []byte(`
{
    "declaredType": "org.cuahsi.waterml.TimeSeriesResponseType",
    "globalScope": true,
    "name": "ns1:timeSeriesResponseType",
    "nil": false,
    "scope": "javax.xml.bind.JAXBElement$GlobalScope",
    "typeSubstituted": false,
    "value": {
        "queryInfo": {
            "criteria": {
                "locationParam": "[ALL:01646500]",
                "parameter": [],
                "variableParam": "[00065, 00060, 00010]"
            },
            "note": [
                {
                    "title": "filter:sites",
                    "value": "[ALL:01646500]"
                },
                {
                    "title": "filter:timeRange",
                    "value": "[mode=LATEST, modifiedSince=null]"
                },
                {
                    "title": "filter:methodId",
                    "value": "methodIds=[ALL]"
                },
                {
                    "title": "requestDT",
                    "value": "2024-06-30T14:39:44.952Z"
                },
                {
                    "title": "requestId",
                    "value": "97d99470-36ee-11ef-8171-2cea7f58f5ca"
                },
                {
                    "title": "disclaimer",
                    "value": "Provisional data are subject to revision. Go to http://waterdata.usgs.gov/nwis/help/?provisional for more information."
                },
                {
                    "title": "server",
                    "value": "vaas01"
                }
            ],
            "queryURL": "http://waterservices.usgs.gov/nwis/iv/format=json&parameterCd=00065,00060,00010&sites=01646500"
        },
        "timeSeries": [
            {
                "name": "USGS:01646500:00010:00000",
                "sourceInfo": {
                    "geoLocation": {
                        "geogLocation": {
                            "latitude": 38.94977778,
                            "longitude": -77.12763889,
                            "srs": "EPSG:4326"
                        },
                        "localSiteXY": []
                    },
                    "note": [],
                    "siteCode": [
                        {
                            "agencyCode": "USGS",
                            "network": "NWIS",
                            "value": "01646500"
                        }
                    ],
                    "siteName": "POTOMAC RIVER NEAR WASH, DC LITTLE FALLS PUMP STA",
                    "siteProperty": [
                        {
                            "name": "siteTypeCd",
                            "value": "ST"
                        },
                        {
                            "name": "hucCd",
                            "value": "02070008"
                        },
                        {
                            "name": "stateCd",
                            "value": "24"
                        },
                        {
                            "name": "countyCd",
                            "value": "24031"
                        }
                    ],
                    "siteType": [],
                    "timeZoneInfo": {
                        "daylightSavingsTimeZone": {
                            "zoneAbbreviation": "EDT",
                            "zoneOffset": "-04:00"
                        },
                        "defaultTimeZone": {
                            "zoneAbbreviation": "EST",
                            "zoneOffset": "-05:00"
                        },
                        "siteUsesDaylightSavingsTime": true
                    }
                },
                "values": [
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "4.1 ft from riverbed (middle), [Discontinued]",
                                "methodID": 69930
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "A",
                                "qualifierDescription": "Approved for publication -- Processing and review completed.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2019-10-01T12:15:00.000-04:00",
                                "qualifiers": [
                                    "A"
                                ],
                                "value": "24.3"
                            }
                        ]
                    },
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "1.0 ft from riverbed (bottom), [Discontinued]",
                                "methodID": 69931
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "A",
                                "qualifierDescription": "Approved for publication -- Processing and review completed.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2019-10-01T12:15:00.000-04:00",
                                "qualifiers": [
                                    "A"
                                ],
                                "value": "24.0"
                            }
                        ]
                    },
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "7.1 ft from riverbed (top), [Discontinued]",
                                "methodID": 69932
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "A",
                                "qualifierDescription": "Approved for publication -- Processing and review completed.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2019-10-01T12:30:00.000-04:00",
                                "qualifiers": [
                                    "A"
                                ],
                                "value": "23.9"
                            }
                        ]
                    },
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "From multiparameter sonde, [Discontinued]",
                                "methodID": 69942
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "A",
                                "qualifierDescription": "Approved for publication -- Processing and review completed.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2019-05-27T15:00:00.000-04:00",
                                "qualifiers": [
                                    "A"
                                ],
                                "value": "24.8"
                            }
                        ]
                    },
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "From multiparameter sonde",
                                "methodID": 252060
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "P",
                                "qualifierDescription": "Provisional data subject to revision.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2024-06-30T09:45:00.000-04:00",
                                "qualifiers": [
                                    "P"
                                ],
                                "value": "29.1"
                            }
                        ]
                    }
                ],
                "variable": {
                    "noDataValue": -999999.0,
                    "note": [],
                    "oid": "45807042",
                    "options": {
                        "option": [
                            {
                                "name": "Statistic",
                                "optionCode": "00000"
                            }
                        ]
                    },
                    "unit": {
                        "unitCode": "deg C"
                    },
                    "valueType": "Derived Value",
                    "variableCode": [
                        {
                            "default": true,
                            "network": "NWIS",
                            "value": "00010",
                            "variableID": 45807042,
                            "vocabulary": "NWIS:UnitValues"
                        }
                    ],
                    "variableDescription": "Temperature, water, degrees Celsius",
                    "variableName": "Temperature, water, &#176;C",
                    "variableProperty": []
                }
            },
            {
                "name": "USGS:01646500:00060:00000",
                "sourceInfo": {
                    "geoLocation": {
                        "geogLocation": {
                            "latitude": 38.94977778,
                            "longitude": -77.12763889,
                            "srs": "EPSG:4326"
                        },
                        "localSiteXY": []
                    },
                    "note": [],
                    "siteCode": [
                        {
                            "agencyCode": "USGS",
                            "network": "NWIS",
                            "value": "01646500"
                        }
                    ],
                    "siteName": "POTOMAC RIVER NEAR WASH, DC LITTLE FALLS PUMP STA",
                    "siteProperty": [
                        {
                            "name": "siteTypeCd",
                            "value": "ST"
                        },
                        {
                            "name": "hucCd",
                            "value": "02070008"
                        },
                        {
                            "name": "stateCd",
                            "value": "24"
                        },
                        {
                            "name": "countyCd",
                            "value": "24031"
                        }
                    ],
                    "siteType": [],
                    "timeZoneInfo": {
                        "daylightSavingsTimeZone": {
                            "zoneAbbreviation": "EDT",
                            "zoneOffset": "-04:00"
                        },
                        "defaultTimeZone": {
                            "zoneAbbreviation": "EST",
                            "zoneOffset": "-05:00"
                        },
                        "siteUsesDaylightSavingsTime": true
                    }
                },
                "values": [
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "",
                                "methodID": 69928
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "P",
                                "qualifierDescription": "Provisional data subject to revision.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2024-06-30T09:45:00.000-04:00",
                                "qualifiers": [
                                    "P"
                                ],
                                "value": "1570"
                            }
                        ]
                    }
                ],
                "variable": {
                    "noDataValue": -999999.0,
                    "note": [],
                    "oid": "45807197",
                    "options": {
                        "option": [
                            {
                                "name": "Statistic",
                                "optionCode": "00000"
                            }
                        ]
                    },
                    "unit": {
                        "unitCode": "ft3/s"
                    },
                    "valueType": "Derived Value",
                    "variableCode": [
                        {
                            "default": true,
                            "network": "NWIS",
                            "value": "00060",
                            "variableID": 45807197,
                            "vocabulary": "NWIS:UnitValues"
                        }
                    ],
                    "variableDescription": "Discharge, cubic feet per second",
                    "variableName": "Streamflow, ft&#179;/s",
                    "variableProperty": []
                }
            },
            {
                "name": "USGS:01646500:00065:00000",
                "sourceInfo": {
                    "geoLocation": {
                        "geogLocation": {
                            "latitude": 38.94977778,
                            "longitude": -77.12763889,
                            "srs": "EPSG:4326"
                        },
                        "localSiteXY": []
                    },
                    "note": [],
                    "siteCode": [
                        {
                            "agencyCode": "USGS",
                            "network": "NWIS",
                            "value": "01646500"
                        }
                    ],
                    "siteName": "POTOMAC RIVER NEAR WASH, DC LITTLE FALLS PUMP STA",
                    "siteProperty": [
                        {
                            "name": "siteTypeCd",
                            "value": "ST"
                        },
                        {
                            "name": "hucCd",
                            "value": "02070008"
                        },
                        {
                            "name": "stateCd",
                            "value": "24"
                        },
                        {
                            "name": "countyCd",
                            "value": "24031"
                        }
                    ],
                    "siteType": [],
                    "timeZoneInfo": {
                        "daylightSavingsTimeZone": {
                            "zoneAbbreviation": "EDT",
                            "zoneOffset": "-04:00"
                        },
                        "defaultTimeZone": {
                            "zoneAbbreviation": "EST",
                            "zoneOffset": "-05:00"
                        },
                        "siteUsesDaylightSavingsTime": true
                    }
                },
                "values": [
                    {
                        "censorCode": [],
                        "method": [
                            {
                                "methodDescription": "",
                                "methodID": 69929
                            }
                        ],
                        "offset": [],
                        "qualifier": [
                            {
                                "network": "NWIS",
                                "qualifierCode": "P",
                                "qualifierDescription": "Provisional data subject to revision.",
                                "qualifierID": 0,
                                "vocabulary": "uv_rmk_cd"
                            }
                        ],
                        "qualityControlLevel": [],
                        "sample": [],
                        "source": [],
                        "value": [
                            {
                                "dateTime": "2024-06-30T09:45:00.000-04:00",
                                "qualifiers": [
                                    "P"
                                ],
                                "value": "2.71"
                            }
                        ]
                    }
                ],
                "variable": {
                    "noDataValue": -999999.0,
                    "note": [],
                    "oid": "45807202",
                    "options": {
                        "option": [
                            {
                                "name": "Statistic",
                                "optionCode": "00000"
                            }
                        ]
                    },
                    "unit": {
                        "unitCode": "ft"
                    },
                    "valueType": "Derived Value",
                    "variableCode": [
                        {
                            "default": true,
                            "network": "NWIS",
                            "value": "00065",
                            "variableID": 45807202,
                            "vocabulary": "NWIS:UnitValues"
                        }
                    ],
                    "variableDescription": "Gage height, feet",
                    "variableName": "Gage height, ft",
                    "variableProperty": []
                }
            }
        ]
    }
}
`)

func TestGetName(t *testing.T) {
	name := "POTOMAC RIVER NEAR WASH, DC LITTLE FALLS PUMP STA"
	want := regexp.MustCompile(`\b` + name + `\b`)
	var ts NWISResponse
	err := json.Unmarshal(nwisJSON, &ts)
	outName := ts.SiteName()
	if !want.MatchString(outName) || err != nil {
		t.Fatalf(`Site Name %q, %v, want match for %#q, nil`, outName, err, want)
	}
}

func TestGetSingleSetOfGauges(t *testing.T) {
	var ts NWISResponse
	err := json.Unmarshal(nwisJSON, &ts)
	if err != nil {
		t.Fatalf("Unmarshalling nwis failed: %v", err)
	}

	count := ts.getNWISGaugeCount()
	if count == 0 {
		t.Fatal("No gauges found!")
	}

	gauges, err := ts.getNWISGauges(0)
	if err != nil {
		t.Fatalf("Failed to get any gauges: %v", err)
	}

	if gauges == nil {
		t.Fatalf("Failed to get any gauge")
	}

	t.Logf("Site %s, Time %s, Temp %2.2f, Height %2.2f, Flow %d", gauges.SiteId, gauges.Time, gauges.Temp, gauges.Stage, gauges.Flow)
}
