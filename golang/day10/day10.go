package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x  int
	y  int
	xv int
	yv int
}

type boundary struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func main() {
	points := readPoints()
	fmt.Println("Time")
	t := findMinAreaTime(points)
	futurePoints := fastForwardPoints(points, t)
	printArray(makeSky(futurePoints))
	fmt.Println(t)
}

func readPoints() []point {
	points := []point{}
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return points
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		re := regexp.MustCompile(`-*\d+`)
		nums := re.FindAllString(scanner.Text(), -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		xv, _ := strconv.Atoi(nums[2])
		yv, _ := strconv.Atoi(nums[3])
		points = append(points, point{x: x, y: y, xv: xv, yv: yv})
	}
	return points
}

func findBoundary(points []point) boundary {
	xMax := points[0].x
	yMax := points[0].y
	xMin := points[0].x
	yMin := points[0].y

	for _, point := range points {
		if point.x > xMax {
			xMax = point.x
		}
		if point.y > yMax {
			yMax = point.y
		}
		if point.x < xMin {
			xMin = point.x
		}
		if point.y < yMin {
			yMin = point.y
		}
	}
	return boundary{xMax: xMax, xMin: xMin, yMax: yMax, yMin: yMin}
}

func fastForwardPoints(points []point, time int) []point {
	newPoints := []point{}
	for _, p := range points {
		newPoints = append(newPoints, point{
			x:  p.x + p.xv*time,
			y:  p.y + p.yv*time,
			xv: p.xv,
			yv: p.yv})
	}
	return newPoints
}

func pointExists(x int, y int, points []point) bool {
	for _, point := range points {
		if point.x == x && point.y == y {
			return true
		}
	}
	return false
}

func makeSky(points []point) [][]string {
	bounds := findBoundary(points)
	sky := make([][]string, 1)
	for y := bounds.yMin; y < bounds.yMax+1; y++ {
		row := make([]string, 1)
		for x := bounds.xMin; x < bounds.xMax+1; x++ {
			if pointExists(x, y, points) {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		sky = append(sky, row)
	}
	return sky
}

func printArray(array [][]string) {
	for _, a := range array {
		for _, n := range a {
			fmt.Print(n)
		}
		fmt.Println("")
	}
}

func findMinAreaTime(points []point) int {
	prev := area(findBoundary(points))
	for i := 0; ; i++ {
		newArea := area(findBoundary(fastForwardPoints(points, i)))
		if newArea > prev {
			return i - 1
		}
		prev = newArea
	}
}

func area(bounds boundary) int {
	return (bounds.xMax - bounds.xMin + 1) * (bounds.yMax - bounds.yMin + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
