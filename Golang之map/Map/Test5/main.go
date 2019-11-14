package main

import(
	"fmt"
	"strings"
)

func main() {
	var s = "How do you do"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}