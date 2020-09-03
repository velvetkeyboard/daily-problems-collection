package main

import "fmt"

func twoSum(arr []int, k int) bool {
	acc := 0
	sums := 0
	for _, number := range arr {
		if number >= k {
			continue
		} else if acc+number <= k {
			acc += number
			sums += 1
		}
		if sums == 2 && acc == k {
			return true
		}
	}
	return false
}

func assertEqualBool(a bool, b bool) {
	if a != b {
		panic(fmt.Sprintf("%d is not equal to %d", a, b))
	}
}

func main() {
	assertEqualBool(twoSum([]int{4, 7, 1, -3, 2}, 5), true)
	assertEqualBool(twoSum([]int{4, 7, 2, -3, 2}, 5), false)
	assertEqualBool(twoSum([]int{7, 4, 1, 4, 1}, 6), false)
	assertEqualBool(twoSum([]int{7, 4, 2, 4, 1}, 6), true)
}
