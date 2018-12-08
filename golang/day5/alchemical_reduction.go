package main

import (
	"bufio"
	"fmt"
	"os"
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

	var chemicals []string

	for scanner.Scan() {
		str := scanner.Text()
		if err != nil {
			fmt.Println(err)
			return
		}
		chemicals = append(chemicals, str)
	}

	units := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	fmt.Println(len(react(chemicals[0], units)))

	//Part 2
	max := 10000000
	//newChemical := ""
	for _, unit := range units {
		unitUpper := strings.ToUpper(unit)
		replacer := strings.NewReplacer(unit, "", unitUpper, "")
		reactedChemical := replacer.Replace(chemicals[0])
		reactedPolymer := react(reactedChemical, units)

		size := len(reactedPolymer)
		if size < max {
			//newChemical = reactedPolymer
			max = size
		}
	}

	fmt.Println(max)

}

func react(chemical string, units []string) string {
	reactedChemical := chemical
	for {
		unitFound := false
		for _, unit := range units {
			unit1 := unit + strings.ToUpper(unit)
			unit2 := strings.ToUpper(unit) + unit
			if strings.Index(reactedChemical, unit1) != -1 || strings.Index(reactedChemical, unit2) != -1 {
				replacer := strings.NewReplacer(unit1, "", unit2, "")
				reactedChemical = replacer.Replace(reactedChemical)
				unitFound = true
			}
		}
		if !unitFound {
			break
		}
	}
	return reactedChemical
}
