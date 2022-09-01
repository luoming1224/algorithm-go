package main

import "fmt"

func dfs3(nums []int, depth int, track []int, used []bool, res *[][]int) {
	if depth == len(nums) {
		t := make([]int, len(track))
		copy(t, track)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			track = append(track, nums[i])
			used[i] = true

			dfs3(nums, depth+1, track, used, res)

			used[i] = false
			track = track[:len(track)-1]
		}
	}
}

func permutations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var result [][]int
	track := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	dfs3(nums, 0, track, used, &result)
	return result
}

func main() {
	nums := []int{0, 1}
	fmt.Println(permutations(nums))
}
