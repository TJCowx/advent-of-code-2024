package main

import (
	"advent-of-code-2024/day01"
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
		fmt.Println("Incorrect format, must be 'day-part' format or just 'day' (1-1 or 1)")
		return
	}

	day, part := parts[0], parts[1]

	switch day {
	case "1":
		day01.Run(&part)
	default:
		fmt.Println("Input hasn't been completed")
	}
}
