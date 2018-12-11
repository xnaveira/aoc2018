package day9

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"strconv"
)

var example bool

func insertMarble(m int, r *ring.Ring) *ring.Ring {
	nr := ring.New(1)
	nr.Value = m

	return r.Next().Link(nr).Prev()
}

func Run(input string) (string, string, error) {

	i, err := ioutil.ReadFile(input)
	if err != nil {
		return "", "", err
	}

	var players, marbles int
	_, err = fmt.Sscanf(string(i), "%d players; last marble is worth %d points", &players, &marbles)
	if err != nil {
		return "", "", err
	}

	r0 := ring.New(1)
	r0.Value = 0

	r1 := ring.New(1)
	r1.Value = 1

	r := r0.Link(r1)
	current := r.Next()

	score := make([]int, players)
	for m := 2; m <= marbles; m++ {
		if m%23 == 0 {
			p := m % players
			v := current
			for i := 0; i < 7; i++ {
				v = v.Prev()
			}
			seventh := v.Value.(int)
			score[p] += m
			score[p] += seventh
			current = v.Next()
			v.Prev().Unlink(1)
		} else {
			current = insertMarble(m, current)
		}

		//fmt.Printf("m:%d p:%d: ", m, current.Value.(int)%players)
		//r.Do(func(v interface{}) {
		//	fmt.Printf(" %d ", v)
		//})
		//fmt.Println("")
	}

	max, winner := 0, 0
	for i, s := range score {
		if s > max {
			max = s
			winner = i
		}
	}

	fmt.Println("The winner is: ", winner)
	result1 := strconv.Itoa(max)

	return result1, "", nil
}
