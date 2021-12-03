package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	s "strings"
)

func parta(in []string) int64 {
	elipson := ""
	gamma := ""
	for i := range strings.Trim(in[0], "\n") {
		zeroes := 0
		ones := 0
		for _, line := range in {
			switch line[i] {
			case '0':
				zeroes++
			case '1':
				ones++
			}
		}
		if zeroes > ones {
			gamma += "0"
			elipson += "1"
		} else {
			gamma += "1"
			elipson += "0"
		}
	}
	fElipson, _ := strconv.ParseInt(elipson, 2, 0)
	fGamma, _ := strconv.ParseInt(gamma, 2, 0)
	return fElipson * fGamma
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
