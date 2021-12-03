package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := readInput("input.txt")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func readInput(path string) []string {
	input, err := os.ReadFile(path)
	check(err)

	entries := strings.Split(strings.Trim(string(input), "\n"), "\n")
	return entries
}

func part1(input []string) int {

	sums := calculateSums(input)
	gammaRate := 0
	epsilonRate := 0

	for i, sum := range sums {
		value := calculateDecimal(len(sums) - i - 1)
		if sum > len(input)/2 {
			gammaRate += 1 * value
			epsilonRate += 0 * value
		} else {
			gammaRate += 0 * value
			epsilonRate += 1 * value
		}
	}
	return gammaRate * epsilonRate
}

func calculateDecimal(position int) int {
	return int(math.Pow(2, float64(position)))
}

func calculateSums(entries []string) [12]int {
	var sums [12]int

	for i := 0; i < len(sums); i++ {
		sums[i] = calculateSumForPos(i, entries)
	}

	return sums
}

func calculateSumForPos(pos int, entries []string) int {
	sum := 0
	for _, entry := range entries {
		value := getValueOnPos(pos, entry)
		sum += value
	}
	return sum
}

func getValueOnPos(pos int, entry string) int {
	parts := strings.Split(entry, "")
	value, _ := strconv.Atoi(parts[pos])
	return value
}

func part2(entries []string) int {

	var mostCommonExtraction IdentifierExtraction = func(entries []string, pos int) int {
		sum := calculateSumForPos(pos, entries)
		if float64(sum) >= float64(len(entries))/2 {
			return 1
		} else {
			return 0
		}
	}

	var leastCommonExtraction IdentifierExtraction = func(entries []string, pos int) int {
		sum := calculateSumForPos(pos, entries)
		if float64(sum) >= float64(len(entries))/2 {
			return 0
		} else {
			return 1
		}
	}
	oxygenGeneratorRating := recursiveFilter(entries, 0, 12, mostCommonExtraction)
	co2ScrubberRating := recursiveFilter(entries, 0, 12, leastCommonExtraction)

	return toDecimal(oxygenGeneratorRating[0]) * toDecimal(co2ScrubberRating[0])
}

func toDecimal(binary string) int {
	dec := 0
	for i, _ := range strings.Split(binary, "") {
		value := getValueOnPos(i, binary)
		dec += value * calculateDecimal(len(binary)-i-1)
	}
	return dec
}

func recursiveFilter(entries []string, pos int, maxPos int, extractIdentifier IdentifierExtraction) []string {
	identifier := extractIdentifier(entries, pos)
	var filtered []string

	for _, entry := range entries {
		if getValueOnPos(pos, entry) == identifier {
			filtered = append(filtered, entry)
		}
	}

	if pos == maxPos || len(filtered) == 1 {
		return filtered
	} else {
		return recursiveFilter(filtered, pos+1, maxPos, extractIdentifier)
	}
}

type IdentifierExtraction func(entries []string, pos int) int
