package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	start := 0
	signMarker := 1
	if s[0] == '-' {
		signMarker = -1
		start++
	} else if s[0] == '+' {
		signMarker = 1
		start++
	}

	var res int
	for i := start; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMarker
		}

		res = res*10 + int(s[i]-'0')

		if res*signMarker <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMarker >= math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * signMarker
}

func main() {
	s := "-91283472332"
	fmt.Println(myAtoi(s))
}
