package main

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"log"
	"math"
	"os"
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

func (o position) paint(gc *draw2dimg.GraphicContext) {
	//fmt.Printf("Painting at %d,%d\n", o.x, o.y)
	gc.BeginPath()
	fx, fy := float64(o.x), float64(o.y)
	gc.MoveTo(fx, fy)
	gc.LineTo(fx+1, fy)
	gc.LineTo(fx+1, fy+1)
	gc.LineTo(fx, fy+1)
	gc.Close()
	gc.FillStroke()
}

func (o position) displace(t position) position {
	o.x += t.x
	o.y += t.y
	return o
}

type points []point

func (c points) searchColumn() bool {
	ccounter := map[int]int{}
	for _, p := range c {
		ccounter[p.p.y]++
	}
	max, candidate := 0, 0
	for k, v := range ccounter {
		if v > max {
			max = v
			candidate = k
		}
	}
	magicColumn := []position{}
	for _, p := range c {
		if p.p.y == candidate {
			magicColumn = append(magicColumn, p.p)
		}
	}
	diff := magicColumn[0].x
	for _, m := range magicColumn[1:] {
		if m.x-diff == 1 {
			continue
		} else {
			return false
		}
	}
	return true
}

func (c points) area() int {
	return (c.max().x - c.min().x) * (c.max().y - c.min().y)
}

func (c *points) max() position {
	maxX, maxY := math.MinInt32, math.MinInt32
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

func (c *points) min() position {
	minX, minY := math.MaxInt32, math.MaxInt32
	for _, p := range *c {
		if p.p.x < minX {
			minX = p.p.x
		}
		if p.p.y < minY {
			minY = p.p.y
		}
	}
	if minX <= minY {
		return position{minX, minX}
	} else {
		return position{minY, minY}
	}
}

func (c points) drawTerm(size position) {
	goterm.Clear()
	//size := c.max()
	//for i := -size.x; i < size.x; i++ {
	for i := 0; i < size.x; i++ {
		//for j := -size.y; j < size.y; j++ {
		for j := 0; j < size.y; j++ {
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

func (c points) plot(size position, frame int) {
	//transform := position{(size.x / 5) * 2, (size.y / 5) * 2}
	//transform := position{(size.x / 10), (size.y / 10)}
	dest := image.NewRGBA(image.Rect(0, 0, 2*size.x, 2*size.y))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetLineWidth(0)

	for i := -size.x; i < size.x; i++ {
		for j := -size.y; j < size.y; j++ {
			tp := position{i, j}
			if tp.in(c) {
				gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
				tp.displace(position{size.x, size.y}).paint(gc)
			} else {
				gc.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
				tp.displace(position{size.x, size.y}).paint(gc)
			}
		}
	}
	draw2dimg.SaveToPngFile(fmt.Sprintf("frame%d.png", frame), dest)
}

//func randomColor() color.RGBA {
//
//	max := 255
//
//	R := rand.Intn(max)
//	G := rand.Intn(max)
//	B := rand.Intn(max)
//
//	return color.RGBA{uint8(R), uint8(G), uint8(B), 0xff}
//
//}

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

	allpoints := points{}
	//i := 0
	for scanner.Scan() {

		var x, y, vx, vy int
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%2d, %2d>", &x, &y, &vx, &vy)
		if err != nil {
			log.Fatal(err)
		}
		allpoints = append(allpoints, point{position{x, y}, velocity{vx, vy}})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//size := allpoints.max()
	type areapoints struct {
		area int
		ps   points
	}
	//areas := []areapoints{}
	for i, _ := range allpoints {
		fmt.Println("Frame ", i)
		//copypoints := make(points, len(allpoints))
		//copy(copypoints, allpoints)
		//areas = append(areas, areapoints{allpoints.area(), copypoints})
		fmt.Println(allpoints.area())
		allpoints.transform()
		//allpoints.plot(size, i)
		//allpoints.drawTerm(size)
		//time.Sleep(500 * time.Millisecond)
		if i == 0 {
			//allpoints.drawTerm(size)
			allpoints.plot(allpoints.max(), 98)
		}
		//allpoints.plot(size, i)

		//found := allpoints.searchColumn()
		//if found {
		//	//allpoints.plot(size, i)
		//	fmt.Println("FOUND")
		//}
	}
	//sort.SliceStable(areas, func(i, j int) bool {
	//	return areas[i].area < areas[j].area
	//})
	//for i := 0; i < 3; i++ {
	//	areas[i].ps.plot(size, 100+i)
	//}
	//allpoints.plot(size)
	fmt.Println("Completed.")
}
