package day10

// https://adventofcode.com/2024/day/10

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
	"fmt"
	"image"
	"strconv"
	"strings"
)

var INPUT_PATH = "day10/input.txt"

func Run(part *string) {
	if part == nil {
		part1(INPUT_PATH)
		part2(INPUT_PATH)
	} else if *part == "1" {
		part1(INPUT_PATH)
	} else if *part == "2" {
		part2(INPUT_PATH)
	} else {
		fmt.Println("INVALID INPUT")
	}
}

type board struct {
	positions  [][]int
	bounds     image.Rectangle
	endPathMap map[image.Point]int
}

func parseInput(path string) board {
	var mapped [][]int
	lines := file_reader.ReadIntoStrArr(path)

	for _, line := range lines {
		var row []int
		chars := strings.Split(line, "")
		for _, char := range chars {
			asInt, _ := strconv.Atoi(char)
			row = append(row, asInt)
		}

		mapped = append(mapped, row)
	}

	bounds := image.Rect(0, 0, len(lines[0]), len(lines))

	return board{
		positions:  mapped,
		bounds:     bounds,
		endPathMap: make(map[image.Point]int),
	}
}

func (b *board) isInBounds(p image.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < b.bounds.Max.X && p.Y < b.bounds.Max.Y
}

func (b *board) processDir(p image.Point, d utils.Direction, target int, visited map[image.Point]bool, part2 bool) int {
	next, _ := utils.GetNextDir(p, d)

	if b.isInBounds(next) && b.positions[next.Y][next.X] == target {
		sum := b.getEndCount(next, d, target, visited, part2)

		b.endPathMap[next] = sum
		return sum
	}

	return 0
}

func (b *board) getEndCount(p image.Point, prevDir utils.Direction, start int, visited map[image.Point]bool, part2 bool) int {
	if !part2 {
		if _, exists := visited[p]; exists {
			return 0
		}

		visited[p] = true
	} else {
		if val, exists := b.endPathMap[p]; exists {
			return val
		}
	}

	if start == 9 {
		b.endPathMap[p] = 1
		return 1
	}

	sum := 0

	if prevDir != utils.Right {
		sum += b.processDir(p, utils.Left, start+1, visited, part2)
	}

	if prevDir != utils.Left {
		sum += b.processDir(p, utils.Right, start+1, visited, part2)
	}

	if prevDir != utils.Up {
		sum += b.processDir(p, utils.Down, start+1, visited, part2)
	}

	if prevDir != utils.Down {
		sum += b.processDir(p, utils.Up, start+1, visited, part2)
	}

	b.endPathMap[p] = sum
	return sum
}

func (b *board) solveP1() int {
	sum := 0

	for y, row := range b.positions {
		for x, val := range row {
			if val == 0 {
				sum += b.getEndCount(
					image.Pt(x, y),
					utils.None,
					0,
					make(map[image.Point]bool),
					false,
				)

			}
		}
	}

	return sum
}

func (b *board) solveP2() int {
	sum := 0

	for y, row := range b.positions {
		for x, val := range row {
			if val == 0 {
				sum += b.getEndCount(
					image.Pt(x, y),
					utils.None,
					0,
					make(map[image.Point]bool),
					true,
				)
			}
		}
	}

	return sum
}
func part1(path string) int {
	fmt.Println("DAY 10 PART 1")

	input := parseInput(path)

	timer := utils.BuildTimer()

	timer.Start()
	res := input.solveP1()
	timer.End()

	fmt.Printf("RESULT: %d | TTC: %s\n", res, timer.TimeLapsed())

	return res
}

func part2(path string) int {
	fmt.Println("DAY 10 PART 2")
	input := parseInput(path)

	timer := utils.BuildTimer()

	timer.Start()
	res := input.solveP2()
	timer.End()

	fmt.Printf("RESULT: %d | TTC: %s\n", res, timer.TimeLapsed())

	return res
}
