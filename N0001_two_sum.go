package main

// TwoSum 两个数相加
func TwoSum(nums []int, target int) []int {
	var idx []int
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i >= j {
				continue
			}
			if target-num1 == num2 {
				idx = append(idx, i)
				idx = append(idx, j)
				return idx
			}
		}
	}
	return idx
}
