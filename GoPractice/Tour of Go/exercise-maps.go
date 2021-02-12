package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	v := strings.Fields(s)
	m := make(map[string]int)
	for _, val := range v {
		_, ok := m[val]
		if ok == true {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
