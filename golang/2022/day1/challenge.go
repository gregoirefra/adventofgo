package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func p1(input []string) (part1 int) {
	part1 = 0
	current := 0
	for _, call := range input {
		if call != "" {
			i, _ := strconv.Atoi(call)
			current += i
			if current > part1 {
				part1 = current
			}
		} else {
			current = 0
		}
	}
	return
}

func p2(input []string) (part2 int) {
	var dic [][]string

}

func Do(input []byte) (part1 string, part2 string) {
	contents := string(input)
	split := strings.Split(contents, "\n")
	part1 = fmt.Sprintf("part1 : %d", p1(split))
	return
}
