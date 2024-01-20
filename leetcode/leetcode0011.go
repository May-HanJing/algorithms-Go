package main

func maxArea(height []int) int {
	var maxValue func(a, b int) int
	maxValue = func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	var minValue func(a, b int) int
	minValue = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	left := 0
	right := len(height) - 1
	for left <= right {
		temp := minValue(height[left], height[right]) * (right - left)
		ans = maxValue(temp, ans)
		if height[left] < height[right] {
			left++
		} else {
			right++
		}
	}
	return ans
}
