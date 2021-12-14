package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splitted := strings.Split(s, " ")
	myMap := make(map[string]int)
	for _, word := range splitted {
		myMap[word] += 1
	}
	return myMap
}

func main() {
	wc.Test(WordCount)
}
