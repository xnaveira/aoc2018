package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type event struct {
	date time.Time
	what string
}

type events []event

func (e events) Len() int {
	return len(e)
}

func (e events) Less(i, j int) bool {
	return e[i].date.Before(e[j].date)
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

var e events
var tLayout string = "2006-01-02 15:04"

func Run(input string) (string, string, error) {

	fmt.Println("This is day4")

	f, err := os.Open(input)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		b, err := parseEvent(line)
		if err != nil {
			return "", "", fmt.Errorf("error parsing events ", err)
		}
		e = append(e, b)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	eSorted := make(events, len(e))
	copy(eSorted, e)
	sort.Sort(eSorted)

	//dayCounter := make(map[time.Time]struct{})
	for _, ee := range eSorted {
		fmt.Printf("%s %s\n", ee.date.String(), ee.what)
		//dayCounter[ee.date] = struct{}{}
	}
	//nDays := len(dayCounter)

	//guards := make(map[string]*minuteTable)
	//parseLog(eSorted, guards)

	m := parseLog2(eSorted)

	totalMinutes := make(map[string]int)

	for id, minutes := range m {
		for _, count := range minutes {
			totalMinutes[id] = totalMinutes[id] + count
		}
	}

	sleepiest := ""
	sminutes := 0
	for key, value := range totalMinutes {
		if value > sminutes {
			sminutes = value
			sleepiest = key
		}
	}

	fmt.Printf("sleepiest is %s with %d minutes\n", sleepiest, sminutes)

	mostSleepMinute := 0
	k := 0
	for key, value := range m[sleepiest] {
		if value > k {
			k = value
			mostSleepMinute = key
		}
	}

	fmt.Printf("the most sleeped minute for %s is: %d with %d times\n ", sleepiest, mostSleepMinute, k)
	//fmt.Println(guards)

	mostSleepedMinute := struct {
		id     string
		minute int
		count  int
	}{
		"",
		0,
		0,
	}

	for id, minutes := range m {
		for m, count := range minutes {
			if count > mostSleepedMinute.count {
				mostSleepedMinute.id = id
				mostSleepedMinute.minute = m
				mostSleepedMinute.count = count

			}
		}
	}

	fmt.Printf("second %v", mostSleepedMinute)

	return "", "", nil

}

func parseEvent(line string) (event, error) {
	l := strings.Split(line, "]")

	d, err := time.Parse(tLayout, l[0][1:])
	if err != nil {
		return event{}, err
	}

	w := l[1]

	return event{date: d, what: w}, nil

}

//type guard struct {
//	beginsShift time.Time
//	fallsAsleep []time.Time
//	wakesUp     []time.Time
//}
//
//type guards map[string]guard

type minuteTable map[time.Time][]bool

func parseLog(log events, guards map[string]*minuteTable) {

	var guardId string
	var begin, end time.Time
	var mt minuteTable
	var mts []bool

	for _, e := range log {
		if strings.Contains(e.what, "begins shift") {
			guardId = strings.Split(e.what, " ")[2]
			mt = make(minuteTable)
			mts = make([]bool, 1440) //minutes in a day
			guards[guardId] = &mt
		} else if strings.Contains(e.what, "falls asleep") {
			begin = e.date
		} else if strings.Contains(e.what, "wakes up") {
			end = e.date
			for t := begin; t.Before(end.Add(-time.Minute)); t = t.Add(time.Minute) {
				//fmt.Printf("minutes for: %s, %s\n", guardId, t.String())
				date, _ := time.Parse("2006-01-02", t.String())
				mts[t.Minute()] = true
				mt[date] = mts
			}
		}
	}

}

//type minutes struct {
//	minute int
//	count int
//}

func parseLog2(log events) map[string]map[int]int {

	var guardId string
	var begin, end time.Time
	//var mt minuteTable
	//var minutes map[int]int

	g := make(map[string]map[int]int)

	for _, e := range log {
		if strings.Contains(e.what, "begins shift") {
			guardId = strings.Split(e.what, " ")[2]
			if g[guardId] == nil {
				g[guardId] = make(map[int]int)
			}
		} else if strings.Contains(e.what, "falls asleep") {
			begin = e.date
		} else if strings.Contains(e.what, "wakes up") {
			end = e.date
			for t := begin; t.Before(end.Add(-time.Minute)); t = t.Add(time.Minute) {
				//fmt.Printf("minutes for: %s, %s\n", guardId, t.String())
				//date, _ := time.Parse("2006-01-02", t.String())
				tt := t.String()
				_ = tt
				g[guardId][t.Minute()]++
			}
		}
	}

	fmt.Println(g["#3167"])
	return g

}
