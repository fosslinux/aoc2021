package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	s "strings"
)

func parta(in []string) (string, string) {
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
	return elipson, gamma
}

func slice2map(s []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, value := range s {
		m[value] = struct{}{}
	}
	return m
}

func map2slice(m map[string]struct{}) []string {
	s := make([]string, 0)
	for k := range m {
		s = append(s, k)
	}
	return s
}

func partb(in []string) (string, string) {
	oxygen := slice2map(in)
	co2 := slice2map(in)
	for i := range in[0] {
		_, tmcom := parta(map2slice(oxygen))
		tlcom, _ := parta(map2slice(co2))
		for _, value := range in {
			if value[i] != tmcom[i] && len(oxygen) != 1 {
				delete(oxygen, value)
			}
			if value[i] != tlcom[i] && len(co2) != 1 {
				delete(co2, value)
			}
		}
	}
	for fOxygen := range oxygen {
		for fCo2 := range co2 {
			return fOxygen, fCo2
		}
	}
	// Will never be reached
	return "", ""
}

func main() {
	bin, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", err))
	}
	in := s.Split(string(bin), "\n")
	in = in[:len(in)-1]

	if os.Args[2] == "a" {
		elipson, gamma := parta(in)
		fElipson, _ := strconv.ParseInt(elipson, 2, 0)
		fGamma, _ := strconv.ParseInt(gamma, 2, 0)
		fmt.Println(fElipson * fGamma)
	} else if os.Args[2] == "b" {
		oxygen, co2 := partb(in)
		fOxygen, _ := strconv.ParseInt(oxygen, 2, 0)
		fCo2, _ := strconv.ParseInt(co2, 2, 0)
		fmt.Println(fOxygen * fCo2)
	}
}
