package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	contents, err := os.ReadFile("programming-task-example-data.log")
	if err != nil {
		log.Fatal(err)
	}

	data := extractLogData(string(contents))

	// task 1 number of unique ip addresses
	uniqueIPs := data.ipAddresses.getUniqueValues()
	fmt.Printf("number of unique IPs: %v", len(uniqueIPs))

	// task 3 top 3 active unique addresses
	top3IPs := data.ipAddresses.getTop3Values()
	fmt.Printf("top 3 IPs: %v", top3IPs)

	// task 2 top 3 most visited URLs
	top3URLs := data.urls.getTop3Values()
	fmt.Printf("top 3 Websites: %v", top3URLs)

}

//func getUnique(input []string) ([]string, map[string]int) {
//	uniqueValuesCount := map[string]int{}
//	var uniqueValues []string
//
//	for _, v := range input {
//		if _, ok := uniqueValuesCount[v]; !ok {
//			uniqueValuesCount[v] = 1
//			uniqueValues = append(uniqueValues, v)
//		} else {
//			uniqueValuesCount[v] += 1
//		}
//	}
//
//	return uniqueValues, uniqueValuesCount
//}

//func getTop3Values(input []string) []string {
//	uniqueValues, uniqueValuesCount := getUnique(input)
//
//	sort.SliceStable(uniqueValues, func(i, j int) bool {
//		return uniqueValuesCount[uniqueValues[i]] > uniqueValuesCount[uniqueValues[j]]
//	})
//
//	return uniqueValues[:4]
//}
