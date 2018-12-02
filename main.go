package main

import (
	"fmt"
	"github.com/xnaveira/aoc2018/day1"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		log.Fatal("You need to enter dayX")
	}

	fmt.Println("Welcome to the AoC 2018")
	fmt.Printf("Running %s\n", os.Args[1])

	day1result, day1resultb, err := day1.Run(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The first result for day1 is: %s\n", day1result)
	fmt.Printf("The second result for day1 is: %s\n", day1resultb)

}
