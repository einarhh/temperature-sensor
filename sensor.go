package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var sensorReadings []int16
var currentIndex int

func convertRawToCelsius(raw int16) float64 {
	// 2048 = 0°C, 3000 = 23°C
	celcius := 0.0242*float64(raw) - 49.6

	// Round to 2 decimal places
	return float64(int(celcius*100)) / 100.0
}

func initializeSensor(fileName string) error {
	// Ensure the file exists and is ready for reading
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist: %v", fileName, err)
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", fileName, err)
	}
	defer file.Close()

	// Read the file line by line into the sensorReadings slice
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %v", fileName, err)
		}
		return fmt.Errorf("file %s is empty", fileName)
	}

	for scanner.Scan() {
		var raw int16
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d", &raw)
		if err != nil {
			log.Printf("Error parsing line '%s': %v", line, err)
			continue
		}
		sensorReadings = append(sensorReadings, raw)
	}
	return nil
}

func getTemperature() (float64, error) {
	if currentIndex >= len(sensorReadings) {
		currentIndex = 0 // Reset index if we reach the end of the slice
	}

	raw := sensorReadings[currentIndex]
	currentIndex++

	// Convert the raw value to temperature in Celsius
	temperature := convertRawToCelsius(raw)
	if temperature < -50 || temperature > 50 {
		return 0, fmt.Errorf("temperature out of range: %.2f", temperature)
	}

	// Simulate a delay for reading the temperature
	time.Sleep(100 * time.Millisecond)

	return temperature, nil
}
