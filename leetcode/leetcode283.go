package main

func moveZeroes(nums []int) {
	start := 0
	end := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[start] = nums[i]
			start++
		} else {
			nums[end] = 0
			end--
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}
