package day06

import "testing"

var TEST_PATH = "./test-input.txt"

func TestPart1(t *testing.T) {
	got := part1(TEST_PATH)
	var expected int = 41

	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	got := part2(TEST_PATH)
	var expected int = 6

	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}
