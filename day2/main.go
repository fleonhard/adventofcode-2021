package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type process func(string)

func main() {
	part1()
	part2()
}

func streamFile(path string, process process) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		process(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {

	depth := 0
	position := 0

	streamFile("input.txt", func(line string) {
		parts := strings.Split(line, " ")
		x, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			position += x
			break
		case "up":
			depth -= x
			break
		case "down":
			depth += x
			break
		}
	})

	fmt.Println("Part 1: ", depth*position)
}

func part2() {

	depth := 0
	position := 0
	aim := 0

	streamFile("input.txt", func(line string) {
		parts := strings.Split(line, " ")
		x, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			position += x
			depth += aim * x
			break
		case "up":
			aim -= x
			break
		case "down":
			aim += x
			break
		}
	})

	fmt.Println("Part 2: ", depth*position)
}
