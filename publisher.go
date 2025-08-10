package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const endpointURL = "http://localhost:5000/api/temperature"

// Marshal the measurement data to JSON and send it to the endpoint
func publishMeasurement(measurement TemperatureMeasurement) {
	jsonData, err := json.Marshal(measurement)
	if err != nil {
		log.Printf("Error marshaling measurement to JSON: %v", err)
		return
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
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from server: %s", resp.Status)
		return
	}
	log.Println("Measurement published successfully")
}
