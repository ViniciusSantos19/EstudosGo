package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complemento := target - num

		if index, found := numMap[complemento]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{0}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 7
	array := TwoSum(nums, target)
	fmt.Println(array)
}
