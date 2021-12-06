package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(path string) []string {
	input, err := os.ReadFile(path)
	check(err)

	entries := strings.Split(strings.Trim(string(input), "\n"), "\n")
	return entries
}

func main() {
	input := readInput("input.txt")
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}
