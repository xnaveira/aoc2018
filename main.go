package main

import (
	"fmt"
	"github.com/xnaveira/aoc2018/day1"
	"github.com/xnaveira/aoc2018/day2"
	"github.com/xnaveira/aoc2018/day3"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		log.Fatal("You need to enter dayX")
	}

	fmt.Println("Welcome to the AoC 2018")
	toRun := os.Args[1]
	fmt.Printf("Running %s\n", toRun)

	switch toRun {
	case "day1":
		day1result, day1resultb, err := day1.Run(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The first result for day1 is: %s\n", day1result)
		fmt.Printf("The second result for day1 is: %s\n", day1resultb)
	case "day2":
		day2result, day2resultb, err := day2.Run(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The first result for day2 is: %s\n", day2result)
		fmt.Printf("The second result for day2 is: %s\n", day2resultb)
	case "day3":
		day3result, day3resultb, err := day3.Run(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The first result for day3 is: %s\n", day3result)
		fmt.Printf("The second result for day3 is: %s\n", day3resultb)

	}

}

