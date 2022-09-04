package main

import "fmt"

// https://zhuanlan.zhihu.com/p/390048232
func backtrack4Permutations(nums []int, k int, res *[][]int)  {
	if k == len(nums) {
		t := make([]int, k)
		copy(t, nums)
		*res = append(*res, t)
		return
	}

	for i := k; i < len(nums); i++ {
		nums[k], nums[i] = nums[i], nums[k]
		backtrack4Permutations(nums, k+1, res)
		nums[k], nums[i] = nums[i], nums[k]
	}
}

func permutations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var res [][]int
	backtrack4Permutations(nums, 0, &res)
	return res
}

func main()  {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permutations(nums))
}
