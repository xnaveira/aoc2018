package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const input = 8868

type field [301][301]cell

type cell struct {
	x, y  int
	power int
}

func (f field) get3Power(x, y, z int) int {
	result := 0
	for i := x; i < x+z; i++ {
		for j := y; j < y+z; j++ {
			result += f[i][j].power
		}
	}
	return result
}

func (c *cell) assginPower(s int) {
	x := c.x
	y := c.y
	rack_id := x + 10
	power_level := rack_id * y
	power_level += s
	power_level = power_level * rack_id
	strPowerLevel := strconv.Itoa(power_level)
	var hundreds int
	var err error
	if len(strPowerLevel) < 3 {
		hundreds = 0
	} else {
		strHundreds := strings.Split(strPowerLevel, "")[len(strPowerLevel)-3]
		hundreds, err = strconv.Atoi(strHundreds)
		if err != nil {
			log.Fatal(err)
		}
	}
	level := hundreds - 5
	c.power = level
}

func (f field) print3grid(x, y int) {
	fmt.Printf("%d %d %d\n", f[x][y].power, f[x+1][y].power, f[x+2][y].power)
	fmt.Printf("%d %d %d\n", f[x][y+1].power, f[x+1][y+1].power, f[x+2][y+1].power)
	fmt.Printf("%d %d %d\n", f[x][y+2].power, f[x+1][y+2].power, f[x+2][y+2].power)
}

func main() {

	f := field{}

	for i := 1; i < len(f); i++ {
		for j := 1; j < len(f); j++ {
			f[i][j].x = i
			f[i][j].y = j
			f[i][j].assginPower(input)

		}
	}

	//f.print3grid(21, 61)
	//fmt.Println(f.get3Power(21, 61, 3))
	//find the power square

	max := math.MinInt32
	maxX := 0
	maxY := 0
	maxZ := 0
	size := 3
	for size = 1; size <= 300; size++ {
		for k := 0; k < size; k++ {
			for l := 0; l < size; l++ {
				for m := 1 + k; m < len(f)-size-k; m += size {
					for n := 1 + l; n < len(f)-size-l; n += size {
						//fmt.Println(k, ",", m, ",", n)
						if f.get3Power(m, n, size) >= max {
							max = f.get3Power(m, n, size)
							maxX = m
							maxY = n
							maxZ = size
						}
					}
				}
			}
		}
	}

	fmt.Printf("The result is: %d,%d,%d with power: %d", maxX, maxY, maxZ, max)
}
