package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/logParser"
	"awesomeProject/internal/logger"
	"fmt"
)

func init() {
	config.Init("./internal/config", "config", "yaml")
}

func main() {

	//logs, err := logParser.Parse(config.Values.LogFilePat)
	//if err != nil {
	//	log.Fatal(err)
	//}

	logs, err := logParser.ParseConcurrently(config.Values.LogFilePath, config.Values.BatchWorkers)
	if err != nil {
		logger.Default().Fatal(err)
	}

	//fmt.Printf("number of lines: %v \n", len(logs.Logs))

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
