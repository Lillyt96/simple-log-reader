package main

import (
	"regexp"
	"sort"
)

var (
	ipv4RegexStr = `((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]|[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]|[0-9])`
	urlRegexStr  = `GET (http://example.net)?(.*)`
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

func extractLogData(input string) *logData {
	data := &logData{}
	// clean up unused logs
	cleanRegex := regexp.MustCompile(`( HTTP(.*))`)
	input = cleanRegex.ReplaceAllString(input, "")

	urlRegex := regexp.MustCompile(urlRegexStr)
	urlSubstring := urlRegex.FindAllStringSubmatch(input, -1)

	for _, url := range urlSubstring {
		data.urls.result = append(data.urls.result, url[2])
	}

	ipv4Regex := regexp.MustCompile(ipv4RegexStr)
	data.ipAddresses.result = ipv4Regex.FindAllString(input, -1)

	return data
}

func (r result) getUniqueValues() []string {
	var uniques []string

	uniqueValuesToCount := mapUniqueValuesToCount(r.result)

	for value := range uniqueValuesToCount {
		uniques = append(uniques, value)
	}

	return uniques
}

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
