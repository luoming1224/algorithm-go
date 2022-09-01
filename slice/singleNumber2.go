package main

import "fmt"

func main() {
	nums := []int{5, 5, 3, 5}
	fmt.Println(singleNumber2(nums))

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber2(nums))
}

func singleNumber2(nums []int) int {
	var one, two, three int
	for i := 0; i < len(nums); i++ {
		two |= one & nums[i]
		one ^= nums[i]
		three = one & two

		one &= ^three
		two &= ^three
	}
	return one
}
