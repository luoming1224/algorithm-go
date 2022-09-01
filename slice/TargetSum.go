package main

import (
	"fmt"
	"sort"
)

var res [][]int

func dfs(candidates []int, index, target int, targetSlice []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int, len(targetSlice)) // 重点注意：当前结果加入结果集时，需要复制一份slice
		copy(b, targetSlice)               //
		res = append(res, b)
		return
	}

	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}

		targetSlice = append(targetSlice, candidates[i])
		dfs(candidates, i, target-candidates[i], targetSlice)
		sliceLen := len(targetSlice)
		targetSlice = targetSlice[:sliceLen-1]
	}
}

func findTargetSum(candidates []int, target int) [][]int {
	var targetSlice []int
	if len(candidates) <= 0 {
		return res
	}

	sort.Ints(candidates)

	dfs(candidates, 0, target, targetSlice)
	return res
}

func main() {
	candidates := []int{3, 4, 10, 11}
	target := 10

	res := findTargetSum(candidates, target)
	fmt.Println(res)
}
