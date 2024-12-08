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
	board          [][]string
	height         int
	width          int
	allBeacons     map[image.Point]string
	countAntiNodes int
	nodesByBeacon  map[string][]image.Point
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

			if c == "#" {
				fmt.Println("THERE IS ALREAD YA # IN HERE WHAT THE")
			}

		}
	}

	return BeaconMap{
		board:          board,
		height:         len(lines),
		width:          len(lines[0]),
		allBeacons:     beacons,
		countAntiNodes: 0,
		nodesByBeacon:  nodesByBeacon,
	}
}

func (bm *BeaconMap) IsInBounds(coor image.Point) bool {
	return coor.X >= 0 && coor.X < bm.width && coor.Y >= 0 && coor.Y < bm.height
}

func (bm *BeaconMap) HasBeaconAlready(coord image.Point) bool {
	_, exists := bm.allBeacons[coord]

	return exists
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

				if bm.IsInBounds(fwd) && !bm.HasBeaconAlready(fwd) {
					bm.countAntiNodes += 1
				}

				if bm.IsInBounds(back) && !bm.HasBeaconAlready(back) {
					bm.countAntiNodes += 1
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

	fmt.Printf("RESULT: %d | TIME: %s\n", bMap.countAntiNodes, timer.TimeLapsed())

	return bMap.countAntiNodes
}

func part2(path string) int {
	fmt.Println("DAY 8 PART 2")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
