package main

import "fmt"

func backtrack(nums []int, start int, track []int, res *[][]int) {
	t := make([]int, len(track)) // 重点注意：深拷贝
	copy(t, track)
	*res = append(*res, t)

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func subSet(nums []int) [][]int {
	var result [][]int
	var track []int
	backtrack(nums, 0, track, &result)

	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subSet(nums))
}
