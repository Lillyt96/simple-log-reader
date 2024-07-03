package main

import (
	"awesomeProject/internal/logParser"
	"fmt"
	"log"
)

func main() {
	//logs, err := logParser.Parse("programming-task-example-data.log")
	//if err != nil {
	//	log.Fatal(err)
	//}

	logs, err := logParser.ParseConcurrently("programming-task-example-data.log")
	if err != nil {
		log.Fatal(err)
	}

	// task 1 number of unique ip addresses
	uniqueIPs := logs.FindUniqueIPs()
	fmt.Printf("number of unique IPs: %v", len(uniqueIPs))

	// task 3 top 3 active unique addresses
	top3IPs := logs.FindTopNIPs(3)
	fmt.Printf("\ntop 3 IPs: %+v", top3IPs)

	// task 2 top 3 most visited URLs
	top3URLs := logs.FindTopNUrls(3)
	fmt.Printf("\ntop 3 Websites: %+v", top3URLs)

}
