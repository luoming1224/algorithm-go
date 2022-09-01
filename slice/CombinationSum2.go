package main

import (
	"fmt"
	"sort"
)

func dfs2(candidates []int, start, target int, c []int, res *[][]int) { //结果集res，需要传地址指针，否则无法返回结果，或者直接用全局变量
	if target <= 0 {
		if target == 0 {
			t := make([]int, len(c))
			copy(t, c)
			*res = append(*res, t)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] { // 去除重复的组合
			continue
		}
		if target < candidates[i] {
			continue
		}
		c = append(c, candidates[i])
		dfs2(candidates, i+1, target-candidates[i], c, res) // 重点注意：每个数字在每个组合中只用一次，所以调用dfs时需要使用i+1
		c = c[:len(c)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	var result [][]int
	var c []int
	sort.Ints(candidates)
	dfs2(candidates, 0, target, c, &result)
	return result
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println(combinationSum2(candidates, target))
}
