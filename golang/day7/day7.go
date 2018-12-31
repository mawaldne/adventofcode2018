package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	taskMap := initializeTasks()
	//fmt.Println(taskMap)

	order := taskOrder(taskMap)
	fmt.Println(string(order))
}

func initializeTasks() map[rune][]rune {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem with file: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	taskMap := make(map[rune][]rune)
	taskSet := make(map[rune]bool)

	r, _ := regexp.Compile(`Step .`)
	t, _ := regexp.Compile(`step .`)

	for _, line := range lines[0 : len(lines)-1] {
		requirementStr := strings.Replace(r.FindString(line), "Step ", "", 1)
		taskStr := strings.Replace(t.FindString(line), "step ", "", 1)
		task := rune(taskStr[0])
		requirement := rune(requirementStr[0])

		requirements, ok := taskMap[task]
		if !ok {
			requirements = make([]rune, 0)
		}
		requirements = append(requirements, requirement)
		taskMap[task] = requirements

		//Get name of all tasks
		taskSet[task] = true
		taskSet[requirement] = true
	}

	//Create empty requirements for tasks not in taskMap
	for task, _ := range taskSet {
		_, ok := taskMap[task]
		if !ok {
			taskMap[task] = make([]rune, 0)
		}
	}
	return taskMap
}

func taskOrder(taskMap map[rune][]rune) []rune {
	completedTasks := make([]rune, 0)

	for len(taskMap) > 0 {
		nextTasks := make([]rune, 0)
		for task, requirements := range taskMap {
			if containsAll(completedTasks, requirements) {
				nextTasks = append(nextTasks, task)
			}
		}
		sort.Slice(nextTasks, func(i, j int) bool { return nextTasks[i] < nextTasks[j] })
		completedTasks = append(completedTasks, nextTasks[0])
		delete(taskMap, nextTasks[0])
	}
	return completedTasks
}

func containsAll(a, b []rune) bool {
	for _, i := range b {
		if !contains(a, i) {
			return false
		}
	}
	return true
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
