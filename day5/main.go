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
	return calculateDangerousFields(input, expandHorizontalAndVerticalLines)
}

func part2(input []string) int {

	var complexLineExpander Expand = func(start [2]int, end [2]int) [][2]int {
		return append(
			expandHorizontalAndVerticalLines(start, end),
			expandDiagonalLines(start, end)...,
		)
	}
	return calculateDangerousFields(input, complexLineExpander)

}

func calculateDangerousFields(input []string, expand Expand) int {
	maxX, maxY, points := createPoints(input, expand)

	grid := createGrid(maxX, maxY)

	updateGrid(points, grid)

	return countDangerousFields(grid)
}

func countDangerousFields(grid [][]int) int {
	dangerousFieldCount := 0
	for _, x := range grid {
		for _, value := range x {
			if value >= 2 {
				dangerousFieldCount++
			}
		}
	}

	return dangerousFieldCount
}

func createPoints(input []string, expandSimpleLine Expand) (int, int, [][2]int) {
	maxX := 0
	maxY := 0
	var points [][2]int
	for _, line := range input {
		start, end := convertLine(line)
		if max(start[0], end[0]) > maxX {
			maxX = max(start[0], end[0])
		}
		if max(start[1], end[1]) > maxY {
			maxY = max(start[1], end[1])
		}

		points = append(points, expandSimpleLine(start, end)...)
	}
	return maxX, maxY, points
}

func updateGrid(points [][2]int, grid [][]int) {
	for _, point := range points {
		grid[point[0]][point[1]]++
	}
}

func createGrid(maxX int, maxY int) [][]int {
	var grid [][]int

	for x := 0; x <= maxX; x++ {
		grid = append(grid, []int{})
		for y := 0; y < maxY; y++ {
			grid[x] = make([]int, maxY+1)
		}
	}
	return grid
}

func convertLine(line string) ([2]int, [2]int) {
	points := strings.Split(line, " -> ")
	start := strings.Split(points[0], ",")
	end := strings.Split(points[1], ",")

	startX, err := strconv.Atoi(start[0])
	check(err)

	startY, err := strconv.Atoi(start[1])
	check(err)

	endX, err := strconv.Atoi(end[0])
	check(err)

	endY, err := strconv.Atoi(end[1])
	check(err)

	return [2]int{startX, startY}, [2]int{endX, endY}
}

type Expand func(pos1 [2]int, pos2 [2]int) [][2]int

var expandHorizontalAndVerticalLines Expand = func(start [2]int, end [2]int) [][2]int {
	var points [][2]int
	if start[0] == end[0] {
		for y := min(start[1], end[1]); y <= max(start[1], end[1]); y++ {
			points = append(points, [2]int{start[0], y})
		}
	} else if start[1] == end[1] {
		for x := min(start[0], end[0]); x <= max(start[0], end[0]); x++ {
			points = append(points, [2]int{x, start[1]})
		}
	}
	return points
}

var expandDiagonalLines Expand = func(start [2]int, end [2]int) [][2]int {

	var points [][2]int

	x := start[0]
	y := start[1]

	if x == end[0] || y == end[1] {
		return points
	}

	xUp := x < end[0]
	yUp := y < end[1]

	for ((x <= end[0] && xUp) || (x >= end[0] && !xUp)) && ((y <= end[1] && yUp) || (y >= end[1] && !yUp)) {
		points = append(points, [2]int{x, y})
		if xUp {
			x++
		} else {
			x--
		}

		if yUp {
			y++
		} else {
			y--
		}
	}
	return points
}

func min(v1 int, v2 int) int {
	return int(math.Min(float64(v1), float64(v2)))
}
func max(v1 int, v2 int) int {
	return int(math.Max(float64(v1), float64(v2)))
}
