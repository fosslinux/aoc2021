package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")
	in = in[:len(in)-1]

	count := 0
	for i := range in {
		if i == 0 {
			continue
		}
		var cur, prev int
		if cur, err = strconv.Atoi(in[i]); err != nil {
			panic(err)
		}
		if prev, err = strconv.Atoi(in[i-1]); err != nil {
			panic(err)
		}
		if cur > prev {
			count++
		}
	}

	fmt.Println(count)
}
