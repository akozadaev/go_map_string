package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 4, 44, 9}
	n := DeleteDuplicates(nums)
	fmt.Println(n)
}

func DeleteDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}
