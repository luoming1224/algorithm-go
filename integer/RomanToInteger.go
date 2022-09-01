package main

import "fmt"

/**
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

var valueMap map[uint8]int
var indexMap map[uint8]int

func init() {
	valueMap = make(map[uint8]int)
	indexMap = make(map[uint8]int)

	valueMap['I'] = 1
	valueMap['V'] = 5
	valueMap['X'] = 10
	valueMap['L'] = 50
	valueMap['C'] = 100
	valueMap['D'] = 500
	valueMap['M'] = 1000

	indexMap['I'] = 0
	indexMap['V'] = 1
	indexMap['X'] = 2
	indexMap['L'] = 3
	indexMap['C'] = 4
	indexMap['D'] = 5
	indexMap['M'] = 6
}

func romanToInteger(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < strLen; i++ {
		if i+1 < strLen {
			o1 := indexMap[s[i]]
			o2 := indexMap[s[i+1]]
			if o1 < o2 {
				sum += valueMap[s[i+1]] - valueMap[s[i]]
				i++
			} else {
				sum += valueMap[s[i]]
			}
		} else {
			sum += valueMap[s[i]]
		}
	}

	return sum
}

func romanToIntegerFinally(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < strLen; i++ {
		if i+1 < strLen && valueMap[s[i]] < valueMap[s[i+1]] {
			sum -= valueMap[s[i]]
		} else {
			sum += valueMap[s[i]]
		}
	}

	return sum
}

func main() {
	s := "III"
	fmt.Println(romanToIntegerFinally(s))

	s = "LVIII"
	fmt.Println(romanToIntegerFinally(s))

	s = "MCMXCIV"
	fmt.Println(romanToIntegerFinally(s))
}
