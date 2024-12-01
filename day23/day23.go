package day23

import "fmt"

func Run(part *string) {
    if part == nil {
        part1()
        part2()
    } else if *part == "1" {
        part1()
    } else if *part == "2" {
        part2()
    } else {
        fmt.Println("INVALID INPUT")
    }
}

func part1() {
    fmt.Println("NOT IMPLEMENTED")
}

func part2() {
    fmt.Println("NOT IMPLEMENTED")
}
