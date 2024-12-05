package day05

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"strconv"
	"strings"
)

var INPUT_PATH = "day05/input.txt"

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

func getPageOrderRules(orders string) map[int]map[int]bool {
	lines := strings.Split(orders, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	rules := make(map[int]map[int]bool)

	for _, line := range lines {
		order := strings.Split(line, "|")
		firstOrder, _ := strconv.Atoi(order[0])
		secondOrder, _ := strconv.Atoi(order[1])

		if _, exists := rules[firstOrder]; exists {
			rules[firstOrder][secondOrder] = true
		} else {
			nested := make(map[int]bool)
			nested[secondOrder] = true
			rules[firstOrder] = nested
		}
	}

	return rules
}

func getOrders(orders string) [][]int {
	lines := strings.Split(orders, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var mapped [][]int

	for _, line := range lines {
		var pages []int
		for _, pageStr := range strings.Split(line, ",") {
			page, _ := strconv.Atoi(pageStr)
			pages = append(pages, page)

		}
		mapped = append(mapped, pages)

	}

	return mapped
}

func parseInput(path string) (map[int]map[int]bool, [][]int) {
	content := file_reader.Read(path)
	sections := strings.Split(content, "\n\n")

	pageOrderRules := getPageOrderRules(sections[0])
	orders := getOrders(sections[1])

	return pageOrderRules, orders
}

func part1(path string) int {
	fmt.Println("DAY 05 PART 1")
	rules, orders := parseInput(path)

	sum := 0

	for _, orderSet := range orders {
		isOrdered := true
		fmt.Println(orderSet)
		var prevPrinted []int
		for i := 0; i < len(orderSet); i++ {
			prevPrinted = append(prevPrinted, orderSet[i])
			if i == 0 {
				continue
			}
			// This number exists in a rule
			if rulesSecond, exists := rules[orderSet[i]]; exists {
				// Get the values before we're in it
				for _, page := range prevPrinted {
					_, exists := rulesSecond[page]

					if exists {
						isOrdered = false
						break
					}
				}
			}

			if !isOrdered {
				break
			}
		}

		if isOrdered {
			sum += orderSet[len(orderSet)/2]
		}
	}

	fmt.Printf("RESULT: %d\n", sum)

	return sum
}

func part2(path string) int {
	fmt.Println("DAY 05 PART 2")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
