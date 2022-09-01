package main

import "fmt"

/**
	问题
	https://leetcode.com/problems/longest-substring-without-repeating-characters/
 	Given a string s, find the length of the longest substring without repeating characters.

	Example 1:

	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

	解释
	https://www.code-recipe.com/post/longest-substring-without-repeating-characters
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLen := 0
	start := 0
	characterMap := make(map[uint8]int)
	for end := 0; end < n; end++ {
		if index, exist := characterMap[s[end]]; exist {
			localLen := end - start
			if localLen > maxLen {
				maxLen = localLen
			}

			for i := start; i <= index; i++ {
				delete(characterMap, s[i])
			}

			start = index + 1
		}
		characterMap[s[end]] = end
	}
	if n-start > maxLen {
		maxLen = n - start
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	maxLen := 0
	start := 0
	characterMap := make(map[uint8]int)

	for end := 0; end < len(s); end++ {
		for {
			if _, exist := characterMap[s[end]]; exist {
				delete(characterMap, s[start])
				start++
			} else {
				break
			}
		}
		characterMap[s[end]] = 1
		localLen := end - start + 1
		if localLen > maxLen {
			maxLen = localLen
		}
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int {
	n := len(s)
	characterMap := make(map[uint8]int)
	curLen := 0
	maxLen := 0
	st := 0
	i := 0
	for ; i < n; i++ {
		if index, exist := characterMap[s[i]]; exist {
			if index >= st {
				curLen = i - st
				if curLen > maxLen {
					maxLen = curLen
				}
				st = index + 1
			}
		}
		characterMap[s[i]] = i
	}
	curLen = i - st
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(s))

	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring2(s))

	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring2(s))
}
