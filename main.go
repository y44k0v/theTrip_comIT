package main

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// gloabal variable for ficticious location with latititude and longitude
var (
	cities = map[string][]float64{
		"Rome":       []float64{12.2, 23.6},
		"Barcelona":  []float64{33.2, 53.6},
		"Marrakech":  []float64{22.2, 77.6},
		"Lisbon":     []float64{43.2, 16.6},
		"Florence":   []float64{35.2, 34.6},
		"Casablanca": []float64{67.2, 88.6},
	}
	city1 string
	city2 string
)

// travel calculation for distance and time
func calculateTravelInfo(origin, destination string) (float64, string) {

	// basic catersian distance calculation
	distance := math.Sqrt(math.Pow(cities[destination][0]-cities[origin][0], 2.0) +
		math.Pow(cities[destination][1]-cities[origin][1], 2.0))

	travelTime := ""

	// Estimation of travel time based on distance
	if distance == 0 {

		travelTime = "0 hours, go wander!!"
	} else if distance < 50 {

		travelTime = "3 hours"
	} else if distance >= 50 && distance < 100 {

		travelTime = "5 hours"
	} else if distance >= 100 && distance < 141 {

		travelTime = "8 hours"
	} else {

		travelTime = "Unknown"
	}
	return distance, travelTime
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Trip Calculator")
	myWindow.Resize(fyne.NewSize(300, 200))

	title := widget.NewLabel("Enter locations:")

	// Anonymous functions to extract the name of the cities in a slice
	Cities := func(cities map[string][]float64) []string {
		var keys []string
		for k := range cities {
			keys = append(keys, k)
		}
		return keys
	}
	// variable with cities' names
	travelCities := Cities(cities)

	origin := widget.NewSelect(travelCities,
		func(s string) {
			fmt.Printf("origin: %s\n", s)
			city1 = s
		})
	destination := widget.NewSelect(travelCities,
		func(s string) {
			fmt.Printf("destination: %s\n", s)
			city2 = s
		})

	//  Trip Results window
	w2 := myApp.NewWindow("Trip Results")
	w2.Resize(fyne.NewSize(150, 150))

	//
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Trip Calculator", Widget: title}},
		OnSubmit: func() { // handle form submission
			// calculatinf distance and travel time
			distance, travelTime := calculateTravelInfo(city1, city2)
			// preparing result
			s1 := fmt.Sprintf("From %s to %s:\n", city1, city2)
			s2 := fmt.Sprintf("Distance: %.2f kilometers\n", distance)
			s3 := fmt.Sprintf("Estimated travel time: %s\n", travelTime)

			// puting results together
			results := func(content ...string) string {
				finalContent := ""
				for _, c := range content {

					finalContent += c
				}
				return finalContent
			}
			// filling content in results window
			w2.SetContent(widget.NewLabel(results(s1, s2, s3)))

			w2.Show()
			myWindow.Close()

		},
	}

	// we can also append items
	form.Append("Origin", origin)
	form.Append("Destination", destination)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
