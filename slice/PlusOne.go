package main

import "fmt"

func plusOne(digits []int) []int {
	numLen := len(digits)
	if numLen == 0 {
		return []int{1}
	}

	carry := 1
	for i := numLen - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] > 9 {
			carry = digits[i] / 10
			digits[i] = digits[i] % 10
		} else {
			carry = 0
			break
		}
	}
	if carry == 1 {
		res := make([]int, 0, numLen+1)
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))

	digits = []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))

	digits = []int{9}
	fmt.Println(plusOne(digits))
}
