package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	s "strings"
)

func parta(input []int) int {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	median := input[len(input)/2-1]
	fuel := 0
	for _, start := range input {
		fuel += int(math.Abs(float64(start - median)))
	}
	return fuel
}

func partb(input []int) int {
	sum := 0
	for _, value := range input {
		sum += value
	}
	mean := int(math.Floor(float64(sum) / float64(len(input))))
	fuel := 0
	for _, start := range input {
		dist := int(math.Abs(float64(start - mean)))
		for i := 1; i < dist+1; i++ {
			fuel += i
		}
	}
	return fuel
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")

	sStart := s.Split(in[0], ",")
	start := make([]int, len(sStart))
	for i, num := range sStart {
		start[i], _ = strconv.Atoi(num)
	}
	if os.Args[2] == "a" {
		fmt.Println(parta(start))
	} else if os.Args[2] == "b" {
		fmt.Println(partb(start))
	}
}
