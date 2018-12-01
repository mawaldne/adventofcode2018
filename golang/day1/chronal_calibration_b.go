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
	scanner.Split(bufio.ScanLines)

	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	printSlice(nums)

	sum := 0
	sums := make(map[int]int)
	sums[0] = 0

	for i := 0; ; {
		sum += nums[i]
		_, ok := sums[sum]
		if ok {
			fmt.Println(sum)
			break
		}

		sums[sum] = sum
		if i == len(nums)-1 {
			i = 0
		} else {
			i += 1
		}
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
