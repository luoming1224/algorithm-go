package main

import (
	"fmt"
	"math"
)

func IsOneEditDistance(s string, t string) bool {
	// write your code here
	len1 := len(s)
	len2 := len(t)
	if (len1 == 0 && len2 == 0) || math.Abs(float64(len1-len2)) > 1 {
		return false
	}
	diffNum := 0
	if len1 == len2 {
		for i := 0; i < len1; i++ {
			if s[i] != t[i] {
				diffNum++
				if diffNum > 1 {
					return false
				}
			}
		}
		if diffNum == 0 {
			return false
		} else {
			return true
		}
	} else {
		i := 0
		j := 0
		for i < len1 && j < len2 {
			if s[i] != t[j] {
				diffNum++
				if diffNum > 1 {
					return false
				}
				if len1 > len2 {
					i++
				} else {
					j++
				}
			}
			i++
			j++
		}
	}

	if diffNum == 1 || diffNum == 0 {
		return true
	}
	return false
}

func main() {
	s := "a"
	t := "ab"
	fmt.Println(IsOneEditDistance(s, t))
}
