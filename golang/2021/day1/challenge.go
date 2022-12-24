package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func p1(input []string) (part1 int) {
	past := 0
	for _, element := range input {
		i, _ := strconv.Atoi(element)
		if past < i && past != 0 {
			part1++
		}
		past = i
	}
	return
}

func p2(input []string) (part2 int) {
	past := 0
	a := make([][]string, len(input)-3)
	for i := range a {
		a[i] = input[i : i+3]
	}
	for _, element := range a {
		value := 0
		for _, j := range element {
			p, _ := strconv.Atoi(j)
			value += p
		}
		if past < value && past != 0 {
			part2++
		}
		past = value

	}
	return
}

func Do(input []byte) (part1 string, part2 string) {
	contents := string(input)
	split := strings.Split(contents, "\n")
	part1 = fmt.Sprintf("part1: %d", p1(split))
	part2 = fmt.Sprintf("part2: %d", p2(split))
	return
}
