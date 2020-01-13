package main

import (
	"fmt"
	"jsontostruct/utils"
	"jsontostruct/validation"
	"time"
)

func main() {
	startTime := time.Now()
	hosts := utils.FetchHosts()
	expectedHardwareDetails := utils.ExtractExpectedHardwareDetails()

	//fmt.Println(hosts)
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()

	fmt.Println("Expected Hardware Configuration: ", expectedHardwareDetails)

	validHostList := validation.Valid(hosts, expectedHardwareDetails)
	fmt.Println("Valid hosts ", len(validHostList))
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
