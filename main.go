package main

import (
	"log"
	"os"
	"time"
)

const fileName = "temperature.txt"

type TemperatureMeasurement struct {
	MaxTemperature     float64 `json:"max"`
	MinTemperature     float64 `json:"min"`
	AverageTemperature float64 `json:"avg"`
	Time               struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"time"`
}

func newMeasurement() TemperatureMeasurement {
	return TemperatureMeasurement{
		MaxTemperature:     -1e9, // Initialize to very low value
		MinTemperature:     1e9,  // Initialize to very high value
		AverageTemperature: 0,
		Time: struct {
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
		}{
			Start: time.Now().UTC(),
			End:   time.Time{},
		},
	}
}

func main() {

	// Initialize the sensor by reading the temperature data from a file
	initializeSensor(fileName)

	log.Println("Sensor and buffer is initialized, reading temperature data...")

	var avgPeriod = 2 * time.Minute
	var totalTemperature float64
	var count int

	measurement := newMeasurement()

	for {
		temp, err := getTemperature()
		if err == os.ErrClosed {
			log.Println("No more temperature data available, exiting...")
			break
		} else if err != nil {
			log.Printf("Error reading temperature: %v", err)
			break
		}
		log.Printf("Current temperature: %.2f", temp)

		totalTemperature += temp
		count++
		if temp > measurement.MaxTemperature {
			measurement.MaxTemperature = temp
		}
		if temp < measurement.MinTemperature {
			measurement.MinTemperature = temp
		}
		if time.Since(measurement.Time.Start) >= avgPeriod {
			averageTemperature := totalTemperature / float64(count)
			measurement.AverageTemperature = float64(int(averageTemperature*100)) / 100.0
			measurement.Time.End = time.Now().UTC()

			publishMeasurement(measurement)

			// Reset for the next period
			totalTemperature = 0
			count = 0
			measurement = newMeasurement()
		}
	}
}
