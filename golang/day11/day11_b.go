package main

import (
	"fmt"
	"math"
)

type gridPower struct {
	x              int
	y              int
	totalPowerSize int
	power          int
}

func main() {
	serialNumber := 3463
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}

	// fmt.Println(powerLevel(8, 3, 5))
	// fmt.Println(powerLevel(57, 122, 79))
	// fmt.Println(powerLevel(39, 217, 196))
	// fmt.Println(powerLevel(71, 101, 153))

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y] = powerLevel(serialNumber, x, y)
		}
	}

	maxGridPower := gridPower{x: 0, y: 0, power: 0}
	for totalPowerSize := 0; totalPowerSize < 300; totalPowerSize++ {
		fmt.Println(totalPowerSize)
		for x := 0; x < 300-totalPowerSize; x++ {
			for y := 0; y < 300-totalPowerSize; y++ {
				gp := gridPower{x: x,
					y:              y,
					totalPowerSize: totalPowerSize,
					power:          totalPower(x, y, totalPowerSize, grid)}
				if gp.power > maxGridPower.power {
					maxGridPower = gp
				}
			}
		}
	}
	fmt.Printf("x %d\n", maxGridPower.x)
	fmt.Printf("y %d\n", maxGridPower.y)
	fmt.Printf("totalPowerSize %d\n", maxGridPower.totalPowerSize)
	fmt.Printf("power %d\n", maxGridPower.power)
}

func totalPower(x int, y int, totalPowerSize int, grid [][]int) int {
	total := 0
	for i := x; i < x+totalPowerSize; i++ {
		for j := y; j < y+totalPowerSize; j++ {
			total += grid[i][j]
		}
	}
	return total
}

func powerLevel(serialNumber int, x int, y int) int {
	rack := x + 10
	start := rack * y
	start = start + serialNumber
	start = start * rack
	hundredsDigit := digit(start, 3)
	return hundredsDigit - 5
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
