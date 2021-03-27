package main

import (
	"fmt"
	"strings"
)

func Decode(roman string) int {
	rndMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	scan := strings.Split(roman, "")[0]
	count := rndMap[scan]
	for i, rnd := range strings.Split(roman, "") {
		if i == 0 {
			continue
		}
		if rndMap[scan] >= rndMap[rnd] {
			count += rndMap[rnd]
			scan = rnd
		} else {
			count += (rndMap[rnd] - 2*rndMap[scan])
		}
	}

	return count
}

func main() {
	fmt.Println(Decode("MCMXC"))
}
