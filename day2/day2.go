package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run(input string) (string, string, error) {
	fmt.Println("This is day2")

	f, err := os.Open(input)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	var codes []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		codes = append(codes, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//codes = []string{"abcdef","bababc","abbcde","abcccd","aabcdd","abcdee","ababab"}

	//codes = []string{"abcde","fghij","klmno","pqrst","fguij","axcye","wvxyz"}
	//codes = []string{"fghij","fguij"}


	duplets := 0
	triplets := 0

	for _, c := range codes {
		//if containsDuplicates(c) {
			//fmt.Println("duplicates! ", c)
			duplet, triplet := countMultiples(c)
			duplets = duplets + duplet
			triplets = triplets + triplet
			fmt.Println("Duplets: ", duplets, " ", "Triplets: ", triplets)
			fmt.Println("----")

		//}
	}

	fmt.Printf("Found %d duplets and %d triplets\n", duplets, triplets)
	result1 := duplets * triplets

	var result2 string
	Loop:
		for _, c1 := range codes {
			for _, c2 := range codes {
				k, i := diffCharCount(c1,c2)
				//fmt.Printf("Difference between %s and %s is %d\n", c1,c2,k)
				if k == 1 {
					fmt.Printf("FOUND %s %s at %d\n",c1,c2, i)
					t := strings.Split(c1,"")
					t = append(t[:i],t[i+1:]...)
					result2 = strings.Join(t,"")
					break Loop

				}
			}
		}

		return strconv.Itoa(result1),result2,nil
}

func SortString(s string) string {
	k := strings.Split(s, "")
	sort.Strings(k)
	return strings.Join(k,"")
}

func containsDuplicates(s string) bool {

	ret := false
	fmt.Printf("Searching: %s\n",s)
	scopy := SortString(s)

	//fmt.Println(scopy)

	for i:=0;i<len(scopy)-1;i++ {
		if scopy[i] == scopy[i+1] {
			//fmt.Println("Duplicate found ", scopy," ", i, " ",i+1)
			ret = true
		}
	}
	return ret
}

func countMultiples(s string) (int, int) {

	duplets := 0
	triplets := 0

	scopy := SortString(s)

	fmt.Println(s)
	for i:=0;i<len(scopy)-1;i++ {
		if scopy[i] == scopy[i+1] {
			fmt.Println("Duplicate found ",s, " ", scopy," ", i, " ",i+1)
			if i < len(scopy)-2 {
				if scopy[i+1] == scopy[i+2] {
					fmt.Println("Triplicate found ",s, " ", scopy," ", i, " ",i+1, " ", i+2)
					triplets = 1
					i++
				} else {
					duplets = 1
				}
			} else {
				duplets = 1
			}

		}
	}
	fmt.Println("---")

	return duplets,triplets

}

func diffCharCount(s1, s2 string) (int, int) {
	//s1sorted := SortString(s1)
	//s2sorted := SortString(s2)

	diff := len(s2)
	index := 0

	for i:=0;i<len(s2);i++ {
		//if s1sorted[i] == s2sorted[i] {
		if s1[i] == s2[i] {
			diff--
		} else {
			index = i
		}
	}
	if diff == 1 {
		return diff, index
	} else {
		return diff, 0
	}
}