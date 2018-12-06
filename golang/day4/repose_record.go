package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var records []string

	for scanner.Scan() {
		str := scanner.Text()
		if err != nil {
			fmt.Println(err)
			return
		}
		records = append(records, str)
	}
	sort.Strings(records)

	//Guard id total minutes
	guardTotalSleep := make(map[string]int)

	//Guart id to map of minute counts
	//So map["#10"][39] <- Counts total sleeps on 39th minutes for guard  #10
	guardSleepMinutes := make(map[string]map[int]int)

	id := ""
	start := ""
	end := ""
	for _, record := range records {
		if strings.Contains(record, "Guard") {
			r, _ := regexp.Compile(`#\d*`)
			id = strings.Trim(r.FindString(record), "#")
			sleepMinutes, ok := guardSleepMinutes[id]
			if !ok {
				sleepMinutes = make(map[int]int)
				guardSleepMinutes[id] = sleepMinutes
			}

		} else if strings.Contains(record, "falls") {
			r, _ := regexp.Compile(`:\d*`)
			start = strings.Trim(r.FindString(record), ":")

		} else if strings.Contains(record, "wakes") {
			r, _ := regexp.Compile(`:\d*`)
			end = strings.Trim(r.FindString(record), ":")
		}

		if id != "" && start != "" && end != "" {
			startVal, _ := strconv.Atoi(start)
			endVal, _ := strconv.Atoi(end)
			guardTotalSleep[id] += endVal - startVal

			sleepMinutes := guardSleepMinutes[id]
			for i := startVal; i < endVal; i++ {
				sleepMinutes[i] += 1
			}
			start = ""
			end = ""
		}
	}

	//Part 1
	//Find Max sleep guard
	maxSleep := 0
	maxGuardId := ""
	for k, v := range guardTotalSleep {
		if v > maxSleep {
			maxSleep = v
			maxGuardId = k
		}
	}

	//Find Max guard Minute
	sleepMinutes := guardSleepMinutes[maxGuardId]
	maxSleepMinuteTotal := 0
	maxSleepMinute := -1
	for k, v := range sleepMinutes {
		if v > maxSleepMinuteTotal {
			maxSleepMinuteTotal = v
			maxSleepMinute = k
		}

	}

	fmt.Println(maxSleep, maxGuardId)
	fmt.Println(maxSleepMinuteTotal, maxSleepMinute)
	guardId, _ := strconv.Atoi(maxGuardId)
	fmt.Println(maxSleepMinute * guardId)

	//Part 2
	frequentMinute := 0
	frequentMinuteTotal := 0
	guardIdFrequent := ""
	for guardId, sleepMinutes := range guardSleepMinutes {
		for minute, totalSleepMinutes := range sleepMinutes {
			if totalSleepMinutes > frequentMinuteTotal {
				frequentMinuteTotal = totalSleepMinutes
				frequentMinute = minute
				guardIdFrequent = guardId
			}
		}
	}
	fmt.Println(frequentMinute, frequentMinuteTotal, guardIdFrequent)
	guardIdF, _ := strconv.Atoi(guardIdFrequent)
	fmt.Println(frequentMinute * guardIdF)

}
