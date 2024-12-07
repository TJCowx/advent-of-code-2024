package day07

import "testing"

var TEST_PATH = "./test-input.txt"

func TestPart1(t *testing.T) {
	gotBrute, gotOpt := part1(TEST_PATH)
	var expected int = 3749

	if gotBrute != expected {
		t.Errorf("BRUTE: Got %d, expected %d", gotBrute, expected)
	}

	if gotOpt != expected {
		t.Errorf("OPTIMIZED: Got %d, expected %d", gotOpt, expected)
	}
}

func TestPart2(t *testing.T) {
	gotBrute, gotOpt := part2(TEST_PATH)
	var expected int = 11387

	if gotBrute != expected {
		t.Errorf("Got %d, expected %d", gotBrute, expected)
	}

	if gotOpt != expected {
		t.Errorf("OPTIMIZED: Got %d, expected %d", gotOpt, expected)
	}
}
