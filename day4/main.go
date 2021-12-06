package main

import (
	"fmt"
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

	randomNumbers := strings.Split(input[0], ",")

	boards := readBoards(input)

	for _, randomNumStr := range randomNumbers {

		randomNum, err := strconv.Atoi(randomNumStr)
		check(err)

		for _, board := range boards {

			win, unmarkedSum := checkForWin(board, randomNum)
			if win {
				return unmarkedSum * randomNum
			}
		}

	}
	return 0
}

func checkForWin(board [][]int, randomNum int) (bool, int) {
	var colSums []int
	unmarkedSum := 0
	win := false

	for _, row := range board {

		rowSum := 0

		for colIndex, value := range row {

			if len(colSums) <= colIndex {
				colSums = append(colSums, 0)
			}

			if value == randomNum {
				row[colIndex] = 0
				rowSum += 0
				colSums[colIndex] += 0
			} else {
				unmarkedSum += value
				rowSum += value
				colSums[colIndex] += value
			}
		}
		if rowSum == 0 {
			win = true
		}
	}

	for _, colSum := range colSums {
		if colSum == 0 {
			win = true
			break
		}
	}
	return win, unmarkedSum
}

func readBoards(input []string) [][][]int {
	var boards [][][]int

	boardIndex := -1
	rowIndex := -1
	for i := 1; i < len(input); i++ {
		if strings.Trim(input[i], "") == "" {
			boardIndex++
			rowIndex = -1
			continue
		}

		if len(boards) <= boardIndex {
			boards = append(boards, [][]int{})
		}

		rowIndex++

		if len(boards[boardIndex]) <= rowIndex {
			boards[boardIndex] = append(boards[boardIndex], []int{})
		}

		for _, numAsStr := range strings.Split(input[i], " ") {
			if strings.Trim(numAsStr, "") == "" {
				continue
			}

			num, err := strconv.Atoi(numAsStr)
			check(err)

			boards[boardIndex][rowIndex] = append(boards[boardIndex][rowIndex], num)
		}

	}
	return boards
}

func part2(input []string) int {

	randomNumbers := strings.Split(input[0], ",")

	boards := readBoards(input)
	lastWinRounds := 0
	lastWinScore := 0

	for _, board := range boards {

		rounds, score := playUntilWin(randomNumbers, board)
		if rounds > lastWinRounds {
			lastWinRounds = rounds
			lastWinScore = score
		}

	}
	return lastWinScore
}

func playUntilWin(randomNumbers []string, board [][]int) (int, int) {
	for roundIndex, randomNumStr := range randomNumbers {

		randomNum, err := strconv.Atoi(randomNumStr)
		check(err)

		win, unmarkedSum := checkForWin(board, randomNum)
		if win {
			return roundIndex, unmarkedSum * randomNum
		}
	}
	return -1, -1
}
