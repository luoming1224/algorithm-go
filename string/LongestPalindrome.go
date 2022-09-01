package main

import (
	"fmt"
	"math"
)

func longestPalindromeSubString(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 0
	start := 0
	for i := 0; i < len(s)-1; i++ {
		curMax1 := isPalindromeSubString(s, i, i)
		curMax2 := isPalindromeSubString(s, i, i+1)

		curMax := curMax1
		if curMax2 > curMax {
			curMax = curMax2
		}
		if curMax > maxLen {
			maxLen = curMax
			start = i - (curMax-1)/2
		}
	}
	return s[start : start+maxLen]
}

func isPalindromeSubString(s string, low, high int) int {
	for low >= 0 && high <= len(s)-1 {
		if s[low] != s[high] {
			break
		}
		low--
		high++
	}
	return high - low - 1
}

func isPalindromeSubString2(s []int32, low, high int) int {
	for low >= 0 && high <= len(s)-1 {
		if s[low] != s[high] {
			break
		}
		low--
		high++
	}
	return high - low - 1
}

func longestPalindromeSubString2(s string) string {
	if len(s) <= 1 {
		return s
	}

	var strSlice []int32
	strSlice = append(strSlice, '#')
	for _, c := range s {
		strSlice = append(strSlice, c)
		strSlice = append(strSlice, '#')
	}

	maxLen := 0
	start := 0
	for i := 0; i < len(strSlice); i++ {
		curMax := isPalindromeSubString2(strSlice, i, i)

		if curMax > maxLen {
			maxLen = curMax
			start = i - (curMax)/2
		}
	}

	start = start / 2
	maxLen = maxLen / 2
	return s[start : start+maxLen]
}

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	s := "babad"
	subStr := longestPalindromeSubString2(s)
	fmt.Println(subStr)
}
