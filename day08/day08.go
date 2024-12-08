package day08

// https://adventofcode.com/2024/day/8

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

var INPUT_PATH string = "day08/input.txt"

type AtennaMap struct {
	width       int
	height      int
	beaconNodes map[string][]Coord
	antiNodes   map[Coord]bool
	taken       map[Coord]bool
}

type Coord struct {
	x int
	y int
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

func parseInput(path string) AtennaMap {
	lines := file_reader.ReadIntoStrArr(path)

	nodeMap := make(map[string][]Coord)
	takenNodes := make(map[Coord]bool)
	height := len(lines)
	width := len(lines[0])

	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char != "." {
				coord := Coord{x, y}
				takenNodes[coord] = true
				if val, exists := nodeMap[char]; exists {
					nodeMap[char] = append(val, coord)
				} else {
					var newCoords []Coord
					newCoords = append(newCoords, coord)
					nodeMap[char] = newCoords
				}
			}
		}
	}

	return AtennaMap{
		height:      height,
		width:       width,
		beaconNodes: nodeMap,
		antiNodes:   make(map[Coord]bool),
		taken:       takenNodes,
	}
}

func getExtendedCoords(curr Coord, next Coord) (Coord, Coord) {
	dx := next.x - curr.x
	dy := next.y - curr.y

	extendBack := Coord{
		x: curr.x - dx,
		y: curr.y - dy,
	}

	extendFwd := Coord{
		x: next.x + dx,
		y: next.y + dy,
	}

	return extendBack, extendFwd
}

func (a *AtennaMap) IsInBounds(node Coord) bool {
	return node.x >= 0 && node.y >= 0 && node.x < a.width && node.y < a.height
}

func (a *AtennaMap) IsFree(node Coord) bool {
	if _, exists := a.taken[node]; exists {
		return false
	}

	return true
}

func (a *AtennaMap) Solve() int {
	sum := 0
	for _, nodes := range a.beaconNodes {
		for i, node := range nodes {
			for j := i + 1; j < len(nodes); j++ {
				backNode, fwdNode := getExtendedCoords(node, nodes[j])

				if a.IsInBounds(backNode) && a.IsFree(backNode) {
					sum += 1
				}

				if a.IsInBounds(fwdNode) && a.IsFree(fwdNode) {
					sum += 1
				}
			}
		}

	}

	return sum
}

func part1(path string) int {
	fmt.Println("DAY 8 PART 1")
	atennaMap := parseInput(path)

	timer := utils.BuildTimer()
	timer.Start()

	res := atennaMap.Solve()

	timer.End()

	fmt.Printf("RESULT: %d | TIME: %s\n", res, timer.TimeLapsed())

	return res
}

func part2(path string) int {
	fmt.Println("DAY 8 PART 2")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
