package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	var spinWord string
	for i, word := range strings.Split(str, " ") {
		if len(word) < 5 {
			if len(strings.Split(str, " "))-1 == i {
				spinWord += word
			} else {
				spinWord += (word + " ")
			}
		} else {
			if len(strings.Split(str, " "))-1 == i {
				spinWord += ReverseWord(word)
			} else {
				spinWord += (ReverseWord(word) + " ")
			}
		}
	}

	return spinWord
}

func ReverseWord(str string) (reverse string) {
	for _, i := range str {
		reverse = string(i) + reverse
	}
	return
}

func main() {
	fmt.Println(SpinWords("This is a test"))
}
