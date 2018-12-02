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

	var ids []string
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			chars, sizeDiff := findSimilarCharacters(ids[i], ids[j])
			if sizeDiff == 1 {
				fmt.Println(ids[i])
				fmt.Println(ids[j])
				fmt.Println(chars)
			}
		}
	}
}

func findSimilarCharacters(a string, b string) (string, int) {
	chars := ""
	for i := 0; i < len(a); i++ {
		if []rune(a)[i] == []rune(b)[i] {
			chars = chars + string([]rune(a)[i])
		}
	}
	return chars, len(a) - len(chars)
}
