package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func parta(in []string) int {
	pos := 0
	depth := 0
	for _, instruction := range in {
		parts := s.Split(instruction, " ")
		direction, sMagnitude := parts[0], parts[1]
		magnitude, _ := strconv.Atoi(sMagnitude)
		switch direction {
		case "forward":
			pos += magnitude
		case "down":
			depth += magnitude
		case "up":
			depth -= magnitude
		}
	}
	return pos * depth
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")
	in = in[:len(in)-1]

	if os.Args[2] == "a" {
		fmt.Println(parta(in))
	}
}
