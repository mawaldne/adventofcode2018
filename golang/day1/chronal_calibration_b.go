package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var numbers []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}

	sums := make(map[int]bool)
	sum := 0

	for {
		for _, num := range numbers {
			if _, ok := sums[sum]; ok {
				fmt.Println(sum)
				os.Exit(0)
			} else {
				sums[sum] = true
			}
			sum += num
		}
	}
}
