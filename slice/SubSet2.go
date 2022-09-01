package main

import (
	"fmt"
	"sort"
)

func backtrack2(nums []int, start int, track []int, res *[][]int) {
	t := make([]int, len(track)) // 重点注意：深拷贝   //进入回溯函数，先加入结果集
	copy(t, track)
	*res = append(*res, t)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrack2(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func subSet2(nums []int) [][]int {
	var result [][]int
	var track []int
	sort.Ints(nums)

	backtrack2(nums, 0, track, &result)

	return result
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subSet2(nums))
}
