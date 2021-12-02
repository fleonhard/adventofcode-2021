package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	measurementStrings := strings.Split(string(input), "\n")
	measurements := make([]int, len(measurementStrings))
	prevMeasurement := -1
	increases := 0

	for i, measurementString := range measurementStrings {
		measurements[i], err = strconv.Atoi(measurementString)

		windowMeasurement := 0
		if i >= 2 {
			windowMeasurement = measurements[i-2] + measurements[i-1] + measurements[i]

			if prevMeasurement != -1 && prevMeasurement < windowMeasurement {
				increases++
			}
			prevMeasurement = windowMeasurement
		}
	}

	fmt.Println("Result: ", increases)
}
