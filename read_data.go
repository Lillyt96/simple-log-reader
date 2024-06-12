package main

import (
	"log"
	"os"
	"regexp"
	"sort"
)

var (
	cleanRegexStr = ` HTTP.*`
	urlRegexStr   = `GET (http://example.net)?(.*)`
	ipv4RegexStr  = `((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]|[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]|[0-9])`
)

type logData struct {
	urls        result
	ipAddresses result
}

type result struct {
	result []string
}

type uniqueWithCount struct {
	value string
	count int
}

func readLogFile(filepath string) string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

// extractLogData extracts urls and ips from a log string with the format provided in programming-task-example-data.log.
func extractLogData(input string) logData {
	data := logData{}
	// remove unused information from logs
	cleanRegex := regexp.MustCompile(cleanRegexStr)
	input = cleanRegex.ReplaceAllString(input, "")

	// extract url information
	urlRegex := regexp.MustCompile(urlRegexStr)
	urlSubstring := urlRegex.FindAllStringSubmatch(input, -1)

	for _, url := range urlSubstring {
		data.urls.result = append(data.urls.result, url[2])
	}

	// extract ip information
	ipv4Regex := regexp.MustCompile(ipv4RegexStr)
	data.ipAddresses.result = ipv4Regex.FindAllString(input, -1)

	return data
}

// getUniqueValues returns a list of unique values from a result list.
func (r result) getUniqueValues() []string {
	var uniques []string

	uniqueValuesToCount := mapUniqueValuesToCount(r.result)

	for value := range uniqueValuesToCount {
		uniques = append(uniques, value)
	}

	return uniques
}

// getUniqueValues returns a the top 3 values from a result list.
func (r result) getTop3Values() []uniqueWithCount {
	var uniquesCount []uniqueWithCount

	uniqueCountMap := mapUniqueValuesToCount(r.result)

	for value, count := range uniqueCountMap {
		uniquesCount = append(uniquesCount, uniqueWithCount{
			value: value,
			count: count,
		})
	}

	sort.SliceStable(uniquesCount, func(i, j int) bool {
		return uniquesCount[i].count > uniquesCount[j].count
	})

	if len(uniquesCount) < 3 {
		return uniquesCount
	}

	return uniquesCount[:3]
}

// mapUniqueValuesToCount reads a list and outputs a map of unique values to the count.
func mapUniqueValuesToCount(input []string) map[string]int {
	uniqueCountMap := map[string]int{}

	for _, v := range input {
		if _, ok := uniqueCountMap[v]; !ok {
			uniqueCountMap[v] = 1
		} else {
			uniqueCountMap[v] += 1
		}
	}

	return uniqueCountMap
}
