package main

import (
	"fmt"
	"sort"
	"strings"
)

// LC242
/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := []byte(s)
	tArr := []byte(t)
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})
	sort.Slice(tArr, func(i, j int) bool {
		return tArr[i] < tArr[j]
	})

	if r := strings.Compare(string(sArr), string(tArr)); r == 0 {
		return true
	}

	return false
}


/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 2.8 MB, less than 82.84% of Go online submissions for Valid Anagram.
Next challenges:
Group Anagrams
Palindrome Permutation
Find All Anagrams in a String
Find Resultant Array After Removing Anagrams
*/
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		arr[i] = 0
	}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a'] += 1
	}
	for j := 0; j < len(t); j++ {
		id := t[j]-'a'
		arr[id] -= 1
		if arr[id] < 0 {
			return false
		}
	}

	return true
}

func main()  {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram2(s, t))

	s = "rat"
	t = "car"
	fmt.Println(isAnagram2(s, t))
}
