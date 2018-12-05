package day5

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Run(input string) (string, string, error) {

	rawPolymer, err := ioutil.ReadFile(input)
	if err != nil {
		return "", "", nil
	}

	//rawPolymer := []byte("dabAcCaCBAcCcaDA")
	//rawPolymer := []byte("aAdabAcCaCBAcCcaDA")
	//rawPolymer := []byte("aAdabAcCaCBAcCcaDAbB")
	//rawPolymer := []byte("aAaaAdabAcCaCBAcCcCcCcCcCaDAbB")

	rawPolymer = []byte(strings.TrimSpace(string(rawPolymer)))
	rawPolymer2 := []byte(strings.TrimSpace(string(rawPolymer)))

	//fmt.Println(len(rawPolymer))

	time1 := time.Now()

	rawPolymer = reduce(rawPolymer)

	elapsed1 := time.Since(time1)

	time2 := time.Now()

	alphabet := make(map[string]int)
	for _, c := range rawPolymer {
		alphabet[strings.ToUpper(string(c))] = 0
	}

	//fmt.Println("alphabet ", alphabet)

	var tmpp []byte
	for k, _ := range alphabet {
		//fmt.Println(len(rawPolymer2))
		tmpp = eliminateChar(rawPolymer2, k)
		//fmt.Printf("remove %s: %s\n", k, string(tmpp))
		tmpp = reduce(tmpp)
		//fmt.Printf("reduced: %s\n", string(tmpp))
		alphabet[k] = len(tmpp)
		//fmt.Println(len(tmpp))
	}

	//searchEquals(rawPolymer)

	type cl struct {
		char string
		len  int
	}

	var ss []cl
	for k, v := range alphabet {
		ss = append(ss, cl{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].len < ss[j].len
	})

	elapsed2 := time.Since(time2)

	//fmt.Printf("remove %s gives a len of %d\n", ss[0].char, ss[0].len)
	result1 := fmt.Sprintf("%s, took %s", strconv.Itoa(len(rawPolymer)), elapsed1.String())
	result2 := fmt.Sprintf("%s, took %s", strconv.Itoa(ss[0].len), elapsed2.String())
	return result1, result2, nil

}

func eliminateChar(p []byte, c string) []byte {
	var tp []byte
	for _, pc := range p {
		if strings.ToUpper(string(pc)) != strings.ToUpper(c) {
			tp = append(tp, pc)
		}
	}
	return tp
}

func reduce(p []byte) []byte {
	var lastLength = 0

Loop:
	for {
		for i := 1; i < len(p); i++ {
			if !equalCase(p[i-1], p[i]) {
				//fmt.Printf("eliminating %s %s\n", string(p[i-1]), string(p[i]))
				p = append(p[:i-1], p[i+1:]...)
			}
			//fmt.Println(string(p))
		}
		if len(p) == lastLength {
			break Loop
		} else {
			lastLength = len(p)
		}
	}
	return p
}

func equalCase(a, b byte) bool {

	if strings.ToUpper(string(a)) != strings.ToUpper(string(b)) {
		return true
	}

	var aUpper, bUpper bool

	if strings.ToUpper(string(a)) == string(a) {
		aUpper = true
	} else {
		aUpper = false
	}

	if strings.ToUpper(string(b)) == string(b) {
		bUpper = true
	} else {
		bUpper = false
	}

	return aUpper == bUpper

}

func searchEquals(p []byte) {
	for i := 1; i < len(p); i++ {
		if strings.ToUpper(string(p[i-1])) == strings.ToUpper(string(p[i])) {
			fmt.Printf("found equals at %d, %d: %s\n", i-1, i, string(p[i-3:i+3]))
		}
	}
}
