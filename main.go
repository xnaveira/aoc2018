package main

import (
	"fmt"
	"github.com/xnaveira/aoc2018/day1"
	"github.com/xnaveira/aoc2018/day2"
	"github.com/xnaveira/aoc2018/day3"
	"github.com/xnaveira/aoc2018/day4"
	"github.com/xnaveira/aoc2018/day5"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 1 {
		log.Fatal("You need to enter dayX")
	}

	fmt.Println("Welcome to the AoC 2018")
	toRun := os.Args[1]
	fmt.Printf("Running %s\n", toRun)

	type solution func(string) (string, string, error)

	type result struct {
		result1, result2 string
	}

	type day struct {
		input    string
		solution solution
		result   result
	}

	var days []day

	days = append(days, day{"day1/input.txt", day1.Run, result{"", ""}})
	days = append(days, day{"day2/input.txt", day2.Run, result{"", ""}})
	days = append(days, day{"day3/input.txt", day3.Run, result{"", ""}})
	days = append(days, day{"day4/input.txt", day4.Run, result{"", ""}})
	days = append(days, day{"day5/input.txt", day5.Run, result{"", ""}})

	theday, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	challenge := days[theday-1]
	challenge.result.result1, challenge.result.result2, err = challenge.solution(challenge.input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day %d:\nSolution to the first problem:%s\nSolution to the second problem:%s", theday, challenge.result.result1, challenge.result.result2)

}
