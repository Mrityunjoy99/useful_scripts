package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type EnvironmentVariable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func main() {
	// Input and output file paths
	inputFilePath := "envtoyaml/input.env"
	outputFilePath := "envtoyaml/output.yaml"

	// Read environment variables from the input file
	envSlice, err := readEnvFile(inputFilePath)
	if err != nil {
		fmt.Printf("Error reading environment variables from %s: %v\n", inputFilePath, err)
		return
	}

	// Convert the environment variables to YAML format
	yamlData, err := convertEnvToYAML(envSlice)
	if err != nil {
		fmt.Printf("Error converting environment variables to YAML: %v\n", err)
		return
	}

	// Write the YAML data to the output file
	err = writeYAMLToFile(outputFilePath, yamlData)
	if err != nil {
		fmt.Printf("Error writing YAML data to %s: %v\n", outputFilePath, err)
		return
	}

	fmt.Printf("Environment variables converted to YAML and written to %s\n", outputFilePath)
}

func readEnvFile(filePath string) ([]EnvironmentVariable, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var envSlice []EnvironmentVariable

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env := EnvironmentVariable{
				Name:  parts[0],
				Value: parts[1],
			}
			envSlice = append(envSlice, env)
		}
	}

	return envSlice, nil
}

func convertEnvToYAML(envSlice []EnvironmentVariable) ([]byte, error) {
	yamlData, err := yaml.Marshal(envSlice)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func writeYAMLToFile(filePath string, yamlData []byte) error {
	err := ioutil.WriteFile(filePath, yamlData, 0644)
	return err
}
