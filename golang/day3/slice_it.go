package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type loc struct {
	row int
	col int
}

type lengths struct {
	id      int
	row_len int
	col_len int
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	overlapCount := make(map[loc]int)
	fabricInfo := make(map[loc]lengths)

	for scanner.Scan() {
		params := getParams(`(?P<id>\d*) @ (?P<row>\d*),(?P<col>\d*): (?P<row_len>\d*)x(?P<col_len>\d*)`, scanner.Text())
		id, _ := strconv.Atoi(params["id"])
		row, _ := strconv.Atoi(params["row"])
		col, _ := strconv.Atoi(params["col"])
		row_len, _ := strconv.Atoi(params["row_len"])
		col_len, _ := strconv.Atoi(params["col_len"])
		fabricInfo[loc{row: row, col: col}] = lengths{id: id, row_len: row_len, col_len: col_len}

		for i := 0; i < row_len; i++ {
			for j := 0; j < col_len; j++ {
				overlapCount[loc{row: row + i, col: col + j}]++
			}
		}
	}

	inches := 0
	for _, v := range overlapCount {
		if v > 1 {
			inches++
		}
	}
	fmt.Println(inches)

	//Part b
	for k, v := range fabricInfo {
		row := k.row
		col := k.col

		row_len := v.row_len
		col_len := v.col_len

		overlapped := false
		for i := 0; i < row_len; i++ {
			for j := 0; j < col_len; j++ {
				if overlapCount[loc{row: row + i, col: col + j}] > 1 {
					overlapped = true
				}
			}
		}
		if !overlapped {
			fmt.Printf("%+v %+v\n", k, v)
			break
		}

	}

}

//Thanks stackoverflow.
//https://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-golang-regular-expressions
func getParams(regEx, str string) (paramsMap map[string]string) {
	var compiled = regexp.MustCompile(regEx)
	match := compiled.FindStringSubmatch(str)

	paramsMap = make(map[string]string)
	for i, name := range compiled.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
