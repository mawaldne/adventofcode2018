package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type coordinateDistance struct {
	coordinate int
	distance   int
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var points []point

	var xMin, xMax, yMin, yMax int = 1000, -1, 1000, -1

	for scanner.Scan() {
		str := scanner.Text()
		if err != nil {
			fmt.Println(err)
			return
		}
		numStrings := strings.Split(str, ",")
		x, _ := strconv.Atoi(strings.Trim(numStrings[0], " "))
		y, _ := strconv.Atoi(strings.Trim(numStrings[1], " "))
		p := point{x: x, y: y}
		points = append(points, p)

		xMin = min(xMin, x)
		xMax = max(xMax, x)
		yMin = min(yMin, y)
		yMax = max(yMax, y)

	}

	//Part 1
	counts := make(map[int]int)
	infiniteSet := make(map[int]bool)
	for y := yMin; y < yMax+1; y++ {
		for x := xMin; x < xMax+1; x++ {
			var distances []coordinateDistance
			for i, point := range points {
				d := coordinateDistance{coordinate: i, distance: distance(x, y, point.x, point.y)}
				distances = append(distances, d)
			}
			sort.Slice(distances, func(i, j int) bool {
				return distances[i].distance < distances[j].distance
			})
			if distances[0].distance != distances[1].distance {
				coordinate := distances[0].coordinate
				counts[coordinate] += 1
			}
			if x == xMin || y == yMin || x == xMax || y == yMax {
				infiniteSet[distances[0].coordinate] = true
			}
		}
	}

	//Remove infinite elements
	maxArea := -1
	for k, v := range counts {
		if _, ok := infiniteSet[k]; ok {
			delete(counts, k)
			continue
		}
		if v > maxArea {
			maxArea = v
		}

	}
	fmt.Println(maxArea)

	//Part 2
	count := 0
	for y := yMin; y < yMax+1; y++ {
		for x := xMin; x < xMax+1; x++ {
			sum := 0
			for _, point := range points {
				sum += distance(x, y, point.x, point.y)
			}
			if sum < 10000 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
