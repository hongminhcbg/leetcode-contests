package main

import "fmt"

func calc(l, m, r int) int {
	target := min(l, r) - 1
	if m <= target {
		return 0
	}

	return m - target
}

func evenLower(nums []int) int {
	result := calc(nums[1], nums[0], nums[1])
	for i := 2; i < len(nums); i += 2 {
		r := i+1
		if i+1 == len(nums) {
			r = i-1
		}

		result += calc(nums[i-1], nums[i], nums[r])
	}

	return result
}

func oddLower(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i += 2 {
		r := i+1
		if i+1 == len(nums) {
			r = i-1
		}
		result += calc(nums[i-1], nums[i], nums[r])

	}

	return result
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

func movesToMakeZigzag(nums []int) int {
	el := evenLower(nums)

	ol := oddLower(nums)
	fmt.Println( el, ol)

	return min(el, ol)
}
func main() {
	fmt.Println(movesToMakeZigzag([]int{1,2,1,2,1,1,1}))
}
