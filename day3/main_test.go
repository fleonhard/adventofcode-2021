package main

import "testing"

func TestPart2(t *testing.T) {
	input := readInput("test.txt")
	got := part2(input)
	if got != 230 {
		t.Errorf("part2(input) = %d; want 230", got)
	}
}
