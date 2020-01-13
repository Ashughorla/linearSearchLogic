package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jsontostruct/ironic"
	"os"
)

// FetchHosts Retrive the host information
func FetchHosts() ironic.Data {
	jsonFile, err := os.Open("introspectedData.json")
	if err != nil {
		fmt.Println(err)
	}

	jsonString, _ := ioutil.ReadAll(jsonFile)

	ironicData := ironic.Data{}
	json.Unmarshal([]byte(jsonString), &ironicData)
	return ironicData
}

// ExtractExpectedHardwareDetails it returns the expected hardware details
func ExtractExpectedHardwareDetails() ironic.ExpectedHardwareConfiguration {

	jsonFile, err := os.Open("expectedHardwareConfiguration.json")
	if err != nil {
		fmt.Println(err)
	}

	jsonString, _ := ioutil.ReadAll(jsonFile)

	expectedHardwareConfiguration := ironic.ExpectedHardwareConfiguration{}
	json.Unmarshal([]byte(jsonString), &expectedHardwareConfiguration)
	return expectedHardwareConfiguration
}
