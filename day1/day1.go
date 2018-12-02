package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run(input string) (string, string, error) {
	fmt.Println("This is day1")

	f, err := os.Open(input)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	var changes []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		changes = append(changes,i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//changes = []int{1,-1}
	//changes = []int{3,3,4,-2,-4}
	//changes = []int{-6,3,8,5,-6}
	//changes = []int{7,7,-2,-7,-4}

	initial := 0
	freq := []int{initial}
	for _,c := range changes {
		initial = initial + c
		freq = append(freq,initial)
	}
	//fmt.Println(initial," ",freq)

	output := strconv.Itoa(initial)

	var output2 string
	Loop:
		for  {

			for k,c := range changes {
				initial = initial + c
				freq = append(freq,initial)
				if k == len(changes)-1 {
					if containsDuplicates(freq) {
						fmt.Println("Search the dup")
						i, err := searchFirstDup(freq)
						if err == nil {
							output2 = strconv.Itoa(i)
							break Loop
						}
					}
				}
			}
		}
	if err != nil {
		return "","",err
	}
	return output,output2,nil
}

func searchFirstDup(a []int) (int, error) {

	var temps []int
	type find struct{
		value int
		distance int
	}

	var finds []find
	for i:=0;i<len(a);i++ {
		temps = copyMinusElement(a,i)
		for j:=0;j<len(temps);j++ {
			if a[i] == temps[j] {
				if j-i > 0 {
					finds = append(finds, find{a[i], j})
				}
			}
		}
	}
	if finds != nil {
		sort.SliceStable(finds, func(i,j int) bool {
			return finds[i].distance < finds[j].distance
		})
		return finds[0].value, nil
	}
	return 0, fmt.Errorf("no dup found")
}

func copyMinusElement(slice []int, element int) []int {

	scopy := make([]int,len(slice))
	copy(scopy,slice)

	ret := make([]int,len(scopy)-1)

	copy(ret,append(scopy[:element], scopy[element+1:]...))

	return ret
}

func containsDuplicates(slice []int) bool {

	ret := false

	var results []int


	scopy := make([]int,len(slice))
	copy(scopy,slice)

	sort.SliceStable(scopy, func(i, j int) bool {
		return scopy[i] < scopy[j]
	})
	for i:=0;i<len(scopy)-1;i++ {
		if scopy[i] == scopy[i+1] {
			fmt.Println("Duplicate found ", scopy[i], " ",scopy[i+1])
			results = append(results,scopy[i])
			ret = true
		}
	}
	//fmt.Println("RESULTS ", results)
	return ret
}
