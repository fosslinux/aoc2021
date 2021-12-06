package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func simulateOne(fish uint8, days int) int {
	fishes := []uint8{fish}
	for day := 0; day < days; day++ {
		for i, fish := range fishes {
			if fish == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i]--
			}
		}
		fmt.Printf("Simulated %d days\n", day)
	}
	return len(fishes)
}

func simulate(fishes []uint8, days int) uint64 {
	after := make([]uint64, 6)
	for i := 0; i < 6; i++ {
		after[i] = uint64(simulateOne(uint8(i), days))
	}
	var sum uint64
	sum = 0
	for _, fish := range fishes {
		sum += after[fish]
	}
	return sum
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")

	sStart := s.Split(in[0], ",")
	start := make([]uint8, len(sStart))
	for i, num := range sStart {
		startInt, _ := strconv.Atoi(num)
		start[i] = uint8(startInt)
	}

	if os.Args[2] == "a" {
		fmt.Println(simulate(start, 80))
	} else if os.Args[2] == "b" {
		fmt.Println(simulate(start, 256))
	}
}
