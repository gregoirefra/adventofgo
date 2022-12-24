package main

import (
	"fmt"
	"log"

	"gregoirefra/adventofcode/2022/day1"
	"gregoirefra/adventofcode/common"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	day1p1, day1p2 := day1.Do(common.GetChallenge("2022", "1"))
	fmt.Println("Day 1:")
	fmt.Println(day1p1)
	fmt.Println(day1p2)

	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("emp:", s[1:5])
}
