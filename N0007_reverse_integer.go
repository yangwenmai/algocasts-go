package main

// Reverse 反转整型数
// x = 123, return 321
// x = -123, return -321
// x = 9646324351, return 0
func Reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}
	var Max = 1<<31 - 1
	if x == 0 {
		return 0
	}
	var y int
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > Max {
		y = 0
	}
	if isNeg {
		y = -y
	}
	return y
}
