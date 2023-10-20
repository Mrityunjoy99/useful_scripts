package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Input and output file paths
	inputFilePath := "envtojson/input.env"
	outputFilePath := "envtojson/output.json"

	// Read environment variables from the input file
	envMap, err := readEnvFile(inputFilePath)
	if err != nil {
		fmt.Printf("Error reading environment variables from %s: %v\n", inputFilePath, err)
		return
	}

	// Convert the environment map to JSON
	jsonData, err := convertEnvToJSON(envMap)
	if err != nil {
		fmt.Printf("Error converting environment variables to JSON: %v\n", err)
		return
	}

	// Write the JSON data to the output file
	err = writeJSONToFile(outputFilePath, jsonData)
	if err != nil {
		fmt.Printf("Error writing JSON data to %s: %v\n", outputFilePath, err)
		return
	}

	fmt.Printf("Environment variables converted to JSON and written to %s\n", outputFilePath)
}

func readEnvFile(filePath string) (map[string]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	return envMap, nil
}

func convertEnvToJSON(envMap map[string]string) ([]byte, error) {
	jsonData, err := json.Marshal(envMap)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func writeJSONToFile(filePath string, jsonData []byte) error {
	err := ioutil.WriteFile(filePath, jsonData, 0644)
	return err
}
