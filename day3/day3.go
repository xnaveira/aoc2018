package day3

import (
	"bufio"
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Run(input string) (string, string, error) {

	fmt.Println("This is day3")

	f, err := os.Open(input)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	var boxes boxArray
	//
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		b, err := parseBoxes(line)
		if err != nil {
			return "", "", fmt.Errorf("error parsing boxes ", err)
		}
		boxes = append(boxes, b)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//stringboxes := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	//stringboxes := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4"}
	//stringboxes := []string{"#1 @ 1,3: 4x4"}
	//stringboxes := []string{"#2 @ 3,1: 4x4"}

	//for _, line := range stringboxes {
	//	b, err := parseBoxes(line)
	//	if err != nil {
	//		return "", "", fmt.Errorf("error parsing boxes ", err)
	//	}
	//	boxes = append(boxes, b)
	//
	//}

	var fabric [1000][1000]int
	//var fabric [8][8]int

	//fmt.Println(fabric)
	//
	//fmt.Println("-----")
	//
	markCoord(&fabric, boxes)
	//
	//fmt.Println("-----")
	//
	//fmt.Println(fabric)
	//
	//fmt.Println("-----")
	//
	//paintFabric(fabric)

	result1 := strconv.Itoa(countNonZero(&fabric))
	result2 := searchOnesOnly(&fabric, boxes)

	plotFabric(fabric, boxes)
	//paintFabric(fabric)

	return result1, result2, nil

}

func searchOnesOnly(fabric *[1000][1000]int, boxes boxArray) string {
	//func searchOnesOnly(fabric *[8][8]int, boxes boxArray) string {

	var ret int

	for _, b := range boxes {
		ret = b.id
		//fmt.Printf("Checking box %d\n", ret)
	Loop:
		for i := b.loc.X; i < b.loc.X+b.size.X; i++ {
			for j := b.loc.Y; j < b.loc.Y+b.size.Y; j++ {
				//fmt.Printf(" %d ", fabric[i][j])
				if fabric[i][j] != 1 {
					ret = -1
					break Loop
				}
			}
		}
		if ret != -1 {
			return fmt.Sprintf("#%d", b.id)
		}
	}
	return fmt.Sprintf("#%d", ret)
}

func countNonZero(fabric *[1000][1000]int) int {
	//func countNonZero(fabric *[8][8]int) int {

	lf := len(fabric)
	counter := 0
	for i := 0; i < lf; i++ {
		for j := 0; j < lf; j++ {
			if fabric[i][j] > 1 {
				counter++
			}
		}
	}
	return counter
}

func markCoord(fabric *[1000][1000]int, boxes boxArray) {
	//func markCoord(fabric *[8][8]int, boxes boxArray) {
	for _, b := range boxes {
		for i := b.loc.X; i < b.loc.X+b.size.X; i++ {
			for j := b.loc.Y; j < b.loc.Y+b.size.Y; j++ {
				fabric[i][j]++
			}
		}
	}
}

func paintFabric(fabric [8][8]int) {
	lf := len(fabric)

	for i := 0; i < lf; i++ {
		for j := 0; j <= lf; j++ {
			if j == lf {
				fmt.Print("\n")
			} else {
				fmt.Printf("%d", fabric[i][j])

			}
		}
	}
}

func randomColor() color.RGBA {

	max := 255

	R := rand.Intn(max)
	G := rand.Intn(max)
	B := rand.Intn(max)

	return color.RGBA{uint8(R), uint8(G), uint8(B), 0xff}

}

func plotFabric(fabric [1000][1000]int, boxes boxArray) error {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, len(fabric), len(fabric)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	//gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(0)

	for _, b := range boxes {
		c := randomColor()
		//fmt.Printf("color %v", c)
		if b.id == 275 {
			gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
		}
		gc.SetFillColor(c)
		gc.BeginPath()
		//fmt.Printf("Draw: %d, %d, %d, %d\n", b.loc.X, b.loc.Y, b.size.X, b.size.Y)
		gc.MoveTo(float64(b.loc.Y), float64(b.loc.X)) // Move to a position to start the new path
		gc.LineTo(float64(b.loc.Y+b.size.Y), float64(b.loc.X))
		gc.LineTo(float64(b.loc.Y+b.size.Y), float64(b.loc.X+b.size.X))
		gc.LineTo(float64(b.loc.Y), float64(b.loc.X+b.size.X))
		gc.Close()
		gc.FillStroke()
	}
	// Draw a closed shape

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)

	return nil
}

type coord struct {
	X int
	Y int
}

type box struct {
	id   int
	loc  coord
	size coord
}

type boxArray []box

func parseBoxes(input string) (box, error) {

	splitted := strings.Split(input, " ")

	id, err := strconv.Atoi(splitted[0][1:])
	if err != nil {
		return box{}, err
	}

	locX, err := strconv.Atoi(strings.Split(splitted[2], ",")[0])
	if err != nil {
		return box{}, err
	}

	locY, err := strconv.Atoi(strings.TrimSuffix(strings.Split(splitted[2], ",")[1], ":"))
	if err != nil {
		return box{}, err
	}

	sizeX, err := strconv.Atoi(strings.Split(splitted[3], "x")[0])
	if err != nil {
		return box{}, err
	}

	sizeY, err := strconv.Atoi(strings.Split(splitted[3], "x")[1])
	if err != nil {
		return box{}, err
	}

	retbox := box{
		id: id,
		loc: coord{
			X: locX,
			Y: locY,
		},
		size: coord{
			X: sizeX,
			Y: sizeY,
		},
	}

	return retbox, nil

}
