package main

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"log"
	"os"
	"time"
)

type position struct {
	x, y int
}

type velocity struct {
	vx, vy int
}

type point struct {
	p position
	v velocity
}

func (o *position) in(c points) bool {
	for _, ps := range c {
		if ps.p == *o {
			return true
		}
	}
	return false
}

type points []point

func (c *points) max() position {
	maxX, maxY := 0, 0
	for _, p := range *c {
		if p.p.x > maxX {
			maxX = p.p.x
		}
		if p.p.y > maxY {
			maxY = p.p.y
		}
	}
	if maxX >= maxY {
		return position{maxX, maxX}
	} else {
		return position{maxY, maxY}
	}
}

func (c points) draw(size position) {
	goterm.Clear()
	//size := c.max()
	for i := -size.x; i < size.x; i++ {
		for j := -size.y; j < size.y; j++ {
			tp := position{i, j}
			goterm.MoveCursor(i+size.x, j+size.y)
			if tp.in(c) {
				goterm.Print("#")
			} else {
				goterm.Print(".")
			}
		}
	}
	goterm.Flush()
	fmt.Println("")
}

func (c points) transform() {
	for i, ps := range c {
		ps.p.x += ps.v.vx
		ps.p.y += ps.v.vy
		c[i] = ps
	}
}

func main() {
	fmt.Println("This is day10")

	var f *os.File
	var err error
	f, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	points := points{}
	for scanner.Scan() {

		var x, y, vx, vy int
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%2d, %2d>", &x, &y, &vx, &vy)
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, point{position{x, y}, velocity{vx, vy}})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(points)
	//err = drawPoints(points)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//drawTest()
	size := points.max()
	for _, _ = range points {
		points.draw(size)
		time.Sleep(500 * time.Millisecond)
		points.transform()
	}
	//points.draw()
	fmt.Println("Completed.")
}

func drawPoints(ps []point) error {
	goterm.Clear()
	for _, p := range ps {
		goterm.MoveCursor(p.p.x, p.p.y)
		_, err := goterm.Print(".")
		if err != nil {
			return err
		}
	}
	goterm.Flush()
	return nil
}

func drawTest() {
	x, y := 20, 20
	goterm.Clear()
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			goterm.MoveCursor(i, j)
			goterm.Print("A")
		}
	}
	goterm.Flush()
}
