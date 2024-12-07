package main

import (
	"advent-of-code-2024/day01"
	"advent-of-code-2024/day02"
	"advent-of-code-2024/day03"
	"advent-of-code-2024/day04"
	"advent-of-code-2024/day05"
	"advent-of-code-2024/day06"
	"advent-of-code-2024/day07"
	"advent-of-code-2024/day08"
	"advent-of-code-2024/day09"
	"advent-of-code-2024/day10"
	"advent-of-code-2024/day11"
	"advent-of-code-2024/day12"
	"advent-of-code-2024/day13"
	"advent-of-code-2024/day14"
	"advent-of-code-2024/day15"
	"advent-of-code-2024/day16"
	"advent-of-code-2024/day17"
	"advent-of-code-2024/day18"
	"advent-of-code-2024/day19"
	"advent-of-code-2024/day20"
	"advent-of-code-2024/day21"
	"advent-of-code-2024/day22"
	"advent-of-code-2024/day23"
	"advent-of-code-2024/day24"
	"advent-of-code-2024/day25"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("INVALID INPUT, ONLY 1 ARGUMENT")
	}

	parts := strings.Split(args[1], "-")

	if len(parts) == 0 || len(parts) > 2 {
		log.Fatal("Incorrect format, must be 'day-part' format or just 'day' (1-1 or 1)")
	}

	day := parts[0]
	var part *string
	if len(parts) > 1 {
		part = &parts[1]
	} else {
		part = nil
	}

	dayFuncs := map[string]func(*string){
		"1":  day01.Run,
		"2":  day02.Run,
		"3":  day03.Run,
		"4":  day04.Run,
		"5":  day05.Run,
		"6":  day06.Run,
		"7":  day07.Run,
		"8":  day08.Run,
		"9":  day09.Run,
		"10": day10.Run,
		"11": day11.Run,
		"12": day12.Run,
		"13": day13.Run,
		"14": day14.Run,
		"15": day15.Run,
		"16": day16.Run,
		"17": day17.Run,
		"18": day18.Run,
		"19": day19.Run,
		"20": day20.Run,
		"21": day21.Run,
		"22": day22.Run,
		"23": day23.Run,
		"24": day24.Run,
		"25": day25.Run,
	}

	if runFunc, exists := dayFuncs[day]; exists {
		runFunc(part)
	} else {
		fmt.Printf("Day %s not found", day)
	}
}
