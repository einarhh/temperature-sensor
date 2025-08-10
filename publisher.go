package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const endpointURL = "http://localhost:5000/api/temperature"
const missingEnpointURL = "http://localhost:5000/api/temperature/missing"

var missingMeasurements = make([]TemperatureMeasurement, 0)
var maxMeasurements = 10 // Maximum number of measurements to keep in memory

// Add a missing measurement to the list, removing the oldest one if necessary
func addMissingMeasurement(measurement TemperatureMeasurement) {
	if len(missingMeasurements) >= maxMeasurements {
		// Remove the oldest measurement if we exceed the limit
		missingMeasurements = missingMeasurements[1:]
	}
	missingMeasurements = append(missingMeasurements, measurement)
	log.Printf("Added missing measurement: %+v", measurement)
}

// Marshal the measurement data to JSON and send it to the endpoint
func publishMeasurement(measurement TemperatureMeasurement) {
	jsonData, err := json.Marshal(measurement)
	if err != nil {
		log.Printf("Error marshaling measurement to JSON: %v", err)
		return
	}

	// Check if we have any missing measurements to send
	if len(missingMeasurements) > 0 {
		publishMissingMeasurements()
	}

	log.Printf("Sending measurement to %s: %s", endpointURL, jsonData)

	req, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		addMissingMeasurement(measurement)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from server: %s", resp.Status)
		addMissingMeasurement(measurement)
		return
	}
	log.Println("Measurement published successfully")
}

// Publish any missing measurements to the missing endpoint
func publishMissingMeasurements() {
	jsonData, err := json.Marshal(missingMeasurements)
	if err != nil {
		log.Printf("Error marshaling missing measurements to JSON: %v", err)
		return
	}

	log.Printf("Sending missing measurements to %s: %s", missingEnpointURL, jsonData)

	req, err := http.NewRequest("POST", missingEnpointURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request for missing measurements: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request for missing measurements: %v", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from server for missing measurements: %s", resp.Status)
		return
	}
	log.Println("Missing measurements published successfully")
	missingMeasurements = make([]TemperatureMeasurement, 0) // Clear the list after publishing
}
