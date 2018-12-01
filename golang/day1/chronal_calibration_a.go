package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Todo - read filename from command line?
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	} else {
		defer input.Close()
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}
