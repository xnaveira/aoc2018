package day6

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type place struct {
	X, Y float64
}

func (p place) distance(q place) float64 {
	return math.Abs(p.X-q.X) + math.Abs(p.Y-q.Y)
}

func (p place) isBorder(q place) bool {
	if p.X == 0 || p.X == q.X || p.Y == 0 || p.Y == q.Y {
		return true
	} else {
		return false
	}
}

type places []place

func (p *places) parsePlaces(l string) error {

	x, err := strconv.Atoi(strings.Split(l, ", ")[0])
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(strings.Split(l, ", ")[1])
	if err != nil {
		return err
	}

	*p = append(*p, place{float64(x), float64(y)})

	return nil
}

type areapoint struct {
	ap place
	pl struct {
		q place
		d float64
	}
}

func Run(input string) (string, string, error) {
	fmt.Println("This is day6")

	ps := make(places, 0)

	f, err := os.Open(input)
	const givensize = 10000

	//f, err := os.Open("day6/example.txt")
	//const givensize = 32

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
		err := ps.parsePlaces(line)
		if err != nil {
			return "", "", fmt.Errorf("error parsing places ", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ps)

	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].distance(place{0, 0}) < ps[j].distance(place{0, 0})
	})

	max := place{}

	for _, p := range ps {
		if p.X >= max.X {
			max.X = p.X
		}
		if p.Y >= max.Y {
			max.Y = p.Y
		}
	}

	a := generateArea(max, float64(1000000))

	fmt.Println(ps)

	//Create zones
	for k, areaPoint := range a {
		for _, pplace := range ps {
			dd := pplace.distance(areaPoint.ap)
			if dd < areaPoint.pl.d {
				areaPoint.pl.q = pplace
				areaPoint.pl.d = dd
				a[k] = areaPoint
			} else if dd == areaPoint.pl.d {
				areaPoint.pl.q = place{-1, -1}
				areaPoint.pl.d = dd
				a[k] = areaPoint
			} else {
				continue
			}
		}
	}

	//Get all the zones created
	zones := map[place]places{}
	for _, pos := range a {
		zones[pos.pl.q] = append(zones[pos.pl.q], pos.ap)
	}

	for k, v := range zones {
		for _, p := range v {
			if p.isBorder(max) {
				delete(zones, k)
				break
			}
		}
	}
	delete(zones, place{-1, -1})

	fmt.Println(a)
	fmt.Println(zones)

	count := 0
	winner := place{}
	for k, v := range zones {
		if len(v) >= count {
			count = len(v)
			winner = k
		}
	}

	fmt.Println(count)

	fmt.Printf("And the winner is: %v %d\n", winner, count)
	return1 := strconv.Itoa(count)

	b := generateArea(max, float64(0))

	//Create zones
	for idx, areaPoint := range b {
		for _, pplace := range ps {
			dd := pplace.distance(areaPoint.ap)
			areaPoint.pl.d = areaPoint.pl.d + dd
			b[idx] = areaPoint
		}
	}

	fmt.Println(b)

	safeZone := places{}

	for _, areaPoint := range b {
		if areaPoint.pl.d < givensize {
			safeZone = append(safeZone, areaPoint.ap)
		}
	}

	result2 := strconv.Itoa(len(safeZone))

	return return1, result2, nil
}

func paintArea(z []areapoint) {
	convert := map[place]string{
		{1, 1}:   "a",
		{1, 6}:   "b",
		{8, 3}:   "c",
		{3, 4}:   "d",
		{5, 5}:   "e",
		{8, 9}:   "f",
		{-1, -1}: ".",
	}
	_ = convert
	l := float64(0)
	for _, zz := range z {
		if zz.ap.Y > l {
			fmt.Printf("\n")
			l = zz.ap.Y
		}
		if zz.pl.d == 0 {
			fmt.Printf("%s", strings.ToUpper(convert[zz.pl.q]))
			//fmt.Printf("\t%v\t", zz.ap)
		} else {
			//fmt.Printf("\t%v\t", zz.ap)
			fmt.Printf("%s", convert[zz.pl.q])
		}
	}
	fmt.Printf("\n")

}

func generateArea(fp place, init float64) []areapoint {

	fmt.Printf("Generate area %.0f x %.0f\n", fp.X, fp.Y)
	area := make([]areapoint, 0)

	for i := float64(0); i <= fp.Y; i++ {
		for j := float64(0); j <= fp.X; j++ {
			//for j := fp.Y; j >= float64(0); j-- {
			area = append(area, areapoint{place{j, i}, struct {
				q place
				d float64
			}{q: place{-1, -1}, d: init}})
		}
	}

	return area
}
