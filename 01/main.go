package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func parta(in []int) int {
	count := 0
	for i := range in {
		if i == 0 {
			continue
		}
		if in[i] > in[i-1] {
			count++
		}
	}
	return count
}

func partb(in []int) int {
	count := 0
	for i := range in {
		if i < 3 {
			continue
		}
		a := in[i-3] + in[i-2] + in[i-1]
		b := in[i-2] + in[i-1] + in[i]
		if b > a {
			count++
		}
	}
	return count
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	sin := s.Split(string(bin), "\n")
	sin = sin[:len(sin)-1]
	// Convert into ints
	in := make([]int, 0)
	for _, num := range sin {
		integer, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		in = append(in, integer)
	}

	if os.Args[2] == "a" {
		fmt.Println(parta(in))
	} else if os.Args[2] == "b" {
		fmt.Println(partb(in))
	}
}
