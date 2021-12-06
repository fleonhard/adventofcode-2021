package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 0
	input := readInput("test_input.txt")
	actual := part1(input)
	if actual != expected {
		t.Errorf("part2(input) = %d; want %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	input := readInput("test_input.txt")
	actual := part2(input)
	if actual != expected {
		t.Errorf("part2(input) = %d; want %d", actual, expected)
	}
}
