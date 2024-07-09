package logParser

import "sort"

type uniqueWithCount struct {
	value string
	count int
}

func (l Logs) FindUniqueIPs() []string {
	var ips []string

	for _, log := range l.Logs {
		ips = append(ips, log.Ip)
	}

	return findUnique(ips)
}

func (l Logs) FindTopNIPs(n int) []uniqueWithCount {
	var ips []string

	for _, log := range l.Logs {
		ips = append(ips, log.Ip)
	}

	return findTopN(ips, n)
}

func (l Logs) FindTopNUrls(n int) []uniqueWithCount {
	var ips []string

	for _, log := range l.Logs {
		ips = append(ips, log.URL)
	}

	return findTopN(ips, n)
}

func findUnique(input []string) []string {
	var uniques []string

	uniquesCount := mapUniqueValuesToCount(input)

	for _, v := range uniquesCount {
		uniques = append(uniques, v.value)
	}

	return uniques
}

func findTopN(input []string, n int) []uniqueWithCount {
	uniquesCount := mapUniqueValuesToCount(input)

	sort.SliceStable(uniquesCount, func(i, j int) bool {
		return uniquesCount[i].count > uniquesCount[j].count
	})

	if len(uniquesCount) < n || n == 0 {
		return uniquesCount
	}

	return uniquesCount[:n]
}

// mapUniqueValuesToCount reads a list and outputs a map of unique values to the count.
func mapUniqueValuesToCount(input []string) []uniqueWithCount {
	uniqueCountMap := map[string]int{}

	for _, v := range input {
		if _, ok := uniqueCountMap[v]; !ok {
			uniqueCountMap[v] = 1
		} else {
			uniqueCountMap[v] += 1
		}
	}

	var uniquesCount []uniqueWithCount

	for value, count := range uniqueCountMap {
		uniquesCount = append(uniquesCount, uniqueWithCount{
			value: value,
			count: count,
		})
	}

	return uniquesCount
}
