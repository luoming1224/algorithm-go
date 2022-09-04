package main

import "fmt"

func firstUniqChar(s string) int {
	index := make([]int, 26)
	for i := 0; i < 26; i++ {
		index[i] = 0
	}
	sLen := len(s)
	for i := 0; i <sLen; i++ {
		id := s[i]-'a'
		index[id] += 1
	}
	for i := 0; i <sLen; i++ {
		if index[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	index := make(map[uint8]int, 26)
	sLen := len(s)
	for i := 0; i <sLen; i++ {
		if n, ok := index[s[i]]; !ok {
			index[s[i]] = 1
		} else {
			index[s[i]] = n +1
		}
	}
	for i := 0; i <sLen; i++ {
		if index[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main()  {
	s := "leetcode"
	fmt.Println(firstUniqChar1(s))
}
