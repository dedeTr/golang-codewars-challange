package main

import (
	"fmt"
	"regexp"
	"strings"
)

func solve(str string) int {
	sentence := strings.ToLower(str)
	strRgx := regexp.MustCompile(`[aiueo]`)
	count := 0
	max := 0
	for _, arr := range strRgx.Split(sentence, -1) {
		for _, letter := range strings.Split(arr, "") {
			count += (int([]rune(letter)[0]) - 96)
		}
		if max < count {
			max = count
		}
		count = 0
	}

	return max
}

func main() {
	fmt.Println(solve("zodiacs"))
}
