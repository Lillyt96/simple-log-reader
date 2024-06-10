package main

import (
	"fmt"
)

func main() {
	contents := readLogFile("programming-task-example-data.log")

	logs := extractLogData(contents)

	// task 1 number of unique ip addresses
	uniqueIPs := logs.ipAddresses.getUniqueValues()
	fmt.Printf("number of unique IPs: %v", len(uniqueIPs))

	// task 3 top 3 active unique addresses
	top3IPs := logs.ipAddresses.getTop3Values()
	fmt.Printf("\ntop 3 IPs: %+v", top3IPs)

	// task 2 top 3 most visited URLs
	top3URLs := logs.urls.getTop3Values()
	fmt.Printf("\ntop 3 Websites: %+v", top3URLs)
}
