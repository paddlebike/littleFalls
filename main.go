package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateNWIS(nwis *widget.Label) {
	ts, err := getNWISSite("01646500")
	if err != nil {
		nwis.SetText(fmt.Sprintf("Failed to get site: %v", err))
	}
	nwis.SetText(ts.ToString())
}

func updateObserved(observed *widget.Label) {
	observed.SetText(ObservedToString("brkm2"))
}

func updateForecast(forecast *widget.Label) {
	forecast.SetText(ForecastToString("brkm2"))
}

func updateOWM(owm *widget.Label) {
	owm.SetText(oneCallByGeoPoint(38.94977778, -77.12763889))
}

func main() {
	a := app.New()
	w := a.NewWindow("Little Falls")

	nwis := widget.NewLabel("")
	nwisTab := container.NewTabItem("USGS", nwis)
	updateNWIS(nwis)

	observed := widget.NewLabel("")
	observedTab := container.NewTabItem("Observed", observed)

	forecast := widget.NewLabel("")
	forecastTab := container.NewTabItem("Forecast", forecast)

	owm := widget.NewLabel("")
	owmTab := container.NewTabItem("Weather", owm)

	tabs := container.NewAppTabs(
		nwisTab,
		observedTab,
		forecastTab,
		owmTab,
	)
	tabs.OnSelected = func(tab *container.TabItem) {
		switch tab {
		case nwisTab:
			updateNWIS(nwis)
		case forecastTab:
			updateForecast(forecast)
		case observedTab:
			updateObserved(observed)
		case owmTab:
			updateOWM(owm)
		}
	}
	tabs.Select(nwisTab)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}
