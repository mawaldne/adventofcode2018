package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	threes := 0
	twos := 0
	for scanner.Scan() {
		id := scanner.Text()
		counts := make(map[rune]int)
		for _, rune := range id {
			counts[rune]++
		}

		foundTwo := false
		foundThree := false
		for _, v := range counts {
			if v == 2 && !foundTwo {
				twos++
				foundTwo = true
			}
			if v == 3 && !foundThree {
				threes++
				foundThree = true
			}
		}
	}
	fmt.Println(twos, threes)
	fmt.Println(twos * threes)
}
