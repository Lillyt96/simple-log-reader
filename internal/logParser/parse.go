package logParser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type Log struct {
	Ip     string
	Time   string
	Method string
	URL    string
	Status string
}

type Logs struct {
	logs []Log
}

func Parse(path string) (*Logs, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	var logs []Log
	for _, line := range lines {
		extractedLog := extractData(line)

		if extractedLog != nil {
			logs = append(logs, *extractedLog)
		}
	}

	return &Logs{logs}, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func extractData(logString string) *Log {
	apacheLogRegexStr := "^(\\S*).*\\[(.*)\\]\\s\"(\\S*)\\s(\\S*)\\s([^\"]*)\"\\s(\\S*)\\s(\\S*)\\s\"([^\"]*)\"\\s\"([^\"]*)\"$"

	apacheLogRegex := regexp.MustCompile(apacheLogRegexStr)

	logResults := apacheLogRegex.FindAllStringSubmatch(logString, -1)

	if len(logResults) == 0 {
		fmt.Println("skipping logResult file due to incorrect format: " + logString)

		return nil
	}

	var logResult Log
	for _, result := range logResults {
		logResult.Ip = result[1]
		logResult.Method = result[3]
		logResult.URL = result[4]
		logResult.Time = result[2]
		logResult.Status = result[6]
	}

	return &logResult
}

func ParseConcurrently(path string) (*Logs, error) {
	// buffer channels
	jobs := make(chan string)
	results := make(chan Log)

	// wait group
	wg := new(sync.WaitGroup)

	// open file for reading
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Go over file line by line and queue up jobs
	go func() {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			jobs <- scanner.Text()
		}

		close(jobs)
	}()

	// set up workers and execute jobs
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go extractDataConcurrently(jobs, results, wg)
	}

	// Close the channel when everything was processed
	go func() {
		wg.Wait()
		close(results)
	}()

	// add results
	var logs []Log
	for result := range results {
		logs = append(logs, result)
	}

	return &Logs{logs}, nil
}

func extractDataConcurrently(job <-chan string, results chan<- Log, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	for j := range job {
		apacheLogRegexStr := "^(\\S*).*\\[(.*)\\]\\s\"(\\S*)\\s(\\S*)\\s([^\"]*)\"\\s(\\S*)\\s(\\S*)\\s\"([^\"]*)\"\\s\"([^\"]*)\"$"

		apacheLogRegex := regexp.MustCompile(apacheLogRegexStr)

		logResults := apacheLogRegex.FindAllStringSubmatch(j, -1)

		if len(logResults) == 0 {
			fmt.Println("skipping logResult file due to incorrect format: " + j)
		} else {
			var logResult Log
			for _, result := range logResults {
				logResult.Ip = result[1]
				logResult.Method = result[3]
				logResult.URL = result[4]
				logResult.Time = result[2]
				logResult.Status = result[6]
			}

			results <- logResult
		}

	}
}
