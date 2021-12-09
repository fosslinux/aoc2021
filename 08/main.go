package main

import (
	"fmt"
	"os"
	s "strings"
)

var segCount map[int]int

func parta(outs [][]string) int {
	count := 0
	for _, out := range outs {
		for _, digit := range out {
			for k := range segCount {
				if len(digit) == k {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	segCount = make(map[int]int)
	segCount[2] = 1
	segCount[4] = 4
	segCount[3] = 7
	segCount[7] = 8

	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	sIn := s.Split(string(bin), "\n")
	sIn = sIn[:len(sIn)-1]

	ins := make([]map[string]int, 0)
	outs := make([][]string, 0)
	for _, line := range sIn {
		words := s.Split(line, " ")
		thisIn := make(map[string]int)
		for _, word := range words[:9] {
			thisIn[word] = -1
		}
		ins = append(ins, thisIn)
		thisOut := make([]string, 4)
		for i, word := range words[11:] {
			thisOut[i] = word
		}
		outs = append(outs, thisOut)
	}

	if os.Args[2] == "a" {
		fmt.Println(parta(outs))
	}
}
