# Simple-log-reader
## Overview

This application reads a log file and analyses:
1. The number of unique IP addresses
2. The top 3 most visited URLs 
3. The top 3 most active IP addresses

The results of this analysis is outputted to the terminal. 

## Instructions
### To run application
Upload log file as "programming-task-example-data.log" to root directory

Run file `go run ./...`

See output in terminal.

## To run tests
`go test -v ./...`

## Assumptions
1. All logs will be in the format provided in the example log file. That being of format:
`{ipv4 address} - - [{date/time}] "GET {url} HTTP/1.1" {response code} 3574 {metadata...}`
2. All ip address will be ipv4.
3. All urls in `programming-task-example-data.log` have the parent url of `http://example.net/`
4. For the purpose of analysis, "url" refers to all subpages. I.e. `/faq/how-to-install/` and 
`/faq/how-to/` are treated as separate urls.

