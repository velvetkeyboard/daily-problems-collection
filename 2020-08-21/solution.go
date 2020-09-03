package main

import "fmt"

func singleNumber(arr []int) int {
	counter := make(map[int]int)
	for _, number := range arr {
		if _, ok := counter[number]; ok {
			delete(counter, number)
		} else {
			counter[number] = number
		}
	}
	ret := -1
	for _, ret = range counter {
		break
	}
	return ret
}

func assertEqualInt(a int, b int) {
	if a != b {
		panic(fmt.Sprintf("%d is not equal to %d", a, b))
	}
}

func main() {
	assertEqualInt(singleNumber([]int{4, 3, 2, 4, 1, 3, 2}), 1)
}
