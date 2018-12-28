package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Rule struct {
	pattern []rune
	result  rune
}

type Environment struct {
	plants      []rune
	plantsNext  []rune
	rules       []Rule
	startingPot int
}

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem with file: %v\n", err)
		os.Exit(1)
	}

	environment := initializeEnvironment(string(content))

	//Part 1
	//environment.runRules(20)
	//fmt.Println(environment.countPots())

	//Part 2 - saw that after about generation 162 it grew by a rate of about 73, so
	//Which is - 12203 + (50000000000 - 162) * 73

	last := 0
	for i := 0; i < 1000; i++ {
		environment.runRules(1)
		current := environment.countPots()
		fmt.Printf("Generation: %d, current: %d, last: %d, total: %d\n", i+1, current, last, current-last)
		last = current
	}
}

func initializeEnvironment(content string) Environment {
	lines := strings.Split(content, "\n")
	plantsStr := strings.Replace(lines[0], "initial state: ", "", 1)
	plants := make([]rune, len(plantsStr))
	plantsNext := make([]rune, len(plantsStr))

	for i, plant := range plantsStr {
		plants[i] = plant
		plantsNext[i] = plant
	}

	rules := make([]Rule, len(lines)-3)
	for i, r := range lines[2 : len(lines)-1] {
		ruleStr := strings.Split(r, " => ")
		rules[i] = Rule{pattern: []rune(ruleStr[0]), result: []rune(ruleStr[1])[0]}
	}
	environment := Environment{plants: plants, plantsNext: plantsNext, rules: rules, startingPot: 0}
	return environment
}

func (e *Environment) runRules(generations int) {
	for g := 0; g < generations; g++ {
		e.padPlants()
		for i := 0; i <= len(e.plants)-5; i++ {
			pattern := e.plants[i : i+5]
			for _, rule := range e.rules {
				if Equal(pattern, rule.pattern) {
					e.plantsNext[i+2] = rule.result
					break
				} else {
					e.plantsNext[i+2] = rune('.')
				}
			}
		}
		copy(e.plants, e.plantsNext)
	}
}

func (e *Environment) padPlants() {
	data := []rune{'.', '.', '.', '.', '.'}
	if !Equal(e.plants[0:5], data) {
		e.plants = append(data, e.plants...)
		e.plantsNext = append(data, e.plantsNext...)
		e.startingPot += 5
	}

	if !Equal(e.plants[len(e.plants)-5:len(e.plants)], data) {
		e.plants = append(e.plants, data...)
		e.plantsNext = append(e.plantsNext, data...)
	}
}

func (e *Environment) countPots() int {
	total := 0
	for i, pot := range e.plants {
		if pot == rune('#') {
			total += i - e.startingPot
		}
	}
	return total
}

func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
