package day08

// https://adventofcode.com/2024/day/8

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
	"fmt"
	"image"
	"strings"
)

var INPUT_PATH string = "day08/input.txt"

type BeaconMap struct {
	board         [][]string
	height        int
	width         int
	allBeacons    map[image.Point]string
	antiNodes     map[image.Point]bool
	nodesByBeacon map[string][]image.Point
}

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

func parseInput(path string) BeaconMap {
	lines := file_reader.ReadIntoStrArr(path)

	var board [][]string = make([][]string, len(lines))
	beacons := make(map[image.Point]string)
	nodesByBeacon := make(map[string][]image.Point)

	for y, line := range lines {
		chars := strings.Split(line, "")
		board[y] = make([]string, len(chars))

		for x, c := range chars {
			board[y][x] = c

			if c != "." {
				coord := image.Pt(x, y)
				beacons[coord] = c
				_, exists := nodesByBeacon[c]

				if !exists {
					nodesByBeacon[c] = make([]image.Point, 0)
				}

				nodesByBeacon[c] = append(nodesByBeacon[c], coord)
			}
		}
	}

	return BeaconMap{
		board:         board,
		height:        len(lines),
		width:         len(lines[0]),
		allBeacons:    beacons,
		antiNodes:     make(map[image.Point]bool),
		nodesByBeacon: nodesByBeacon,
	}
}

func (bm *BeaconMap) IsInBounds(coor image.Point) bool {
	return coor.X >= 0 && coor.X < bm.width && coor.Y >= 0 && coor.Y < bm.height
}

func (bm *BeaconMap) PrintSolved() {
	fmt.Println(bm.antiNodes)
	fmt.Println("==========BOARD=============")
	for y, line := range bm.board {
		var out string = ""
		for x := range line {
			coord := image.Point{X: x, Y: y}

			if beacon, exists := bm.allBeacons[coord]; exists {
				out += beacon
			} else if _, exists := bm.antiNodes[coord]; exists {
				out += "#"
			} else {
				out += "."
			}
		}

		fmt.Println(out)
	}
}

func (bm *BeaconMap) Solve() {
	for _, nodes := range bm.nodesByBeacon {
		for i := 0; i < len(nodes); i++ {
			curr := nodes[i]
			for j := i + 1; j < len(nodes); j++ {
				next := nodes[j]

				// Get the vector (delta) between the two nodes
				delta := next.Sub(curr)

				// Calculate the possible anti-nodes by extending both directions
				// Forward: move from curr to next
				fwd := next.Add(delta)
				// Backward: move from next to curr
				back := curr.Sub(delta)

				if bm.IsInBounds(fwd) {
					bm.antiNodes[fwd] = true
				}

				if bm.IsInBounds(back) {
					bm.antiNodes[back] = true
				}
			}
		}
	}
}

func (bm *BeaconMap) SolveP2() {
	for _, nodes := range bm.nodesByBeacon {
		for i := 0; i < len(nodes); i++ {
			curr := nodes[i]
			bm.antiNodes[curr] = true
			for j := i + 1; j < len(nodes); j++ {
				next := nodes[j]

				// Get the vector (delta) between the two nodes
				delta := next.Sub(curr)

				// Calculate the possible anti-nodes by extending both directions
				// Forward: move from curr to next
				fwd := next.Add(delta)

				for bm.IsInBounds(fwd) {
					bm.antiNodes[fwd] = true
					fwd = fwd.Add(delta)
				}

				// Backward: move from next to curr
				back := curr.Sub(delta)

				for bm.IsInBounds(back) {
					bm.antiNodes[back] = true
					back = back.Sub(delta)
				}
			}
		}
	}
}

func part1(path string) int {
	fmt.Println("DAY 8 PART 1")
	bMap := parseInput(path)

	timer := utils.BuildTimer()
	timer.Start()

	bMap.Solve()

	timer.End()

	fmt.Printf("RESULT: %d | TIME: %s\n", len(bMap.antiNodes), timer.TimeLapsed())

	return len(bMap.antiNodes)
}

func part2(path string) int {
	fmt.Println("DAY 8 PART 2")

	bMap := parseInput(path)

	timer := utils.BuildTimer()
	timer.Start()

	bMap.SolveP2()

	timer.End()

	fmt.Printf("RESULT: %d | TIME: %s\n", len(bMap.antiNodes), timer.TimeLapsed())

	return len(bMap.antiNodes)
	return 0
}
