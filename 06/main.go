package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func simulate(fishes []int) int {
	for day := 0; day < 80; day++ {
		for i, fish := range fishes {
			if fish == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i]--
			}
		}
	}
	return len(fishes)
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
	}
}
