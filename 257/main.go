package main

import (
	"fmt"
	"sort"
)

func countQuadruplets(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; j < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						result++
					}
				}
			}

		}
	}

	return result
}

func findFistIndexDiff(arr [][]int) (int, int) {
	n := len(arr)
	maxVal := arr[n-1][1]

	for i := n - 2; i >= 0; i-- {
		if arr[i][0] != arr[n-1][0] {
			return i, maxVal
		}

		if arr[i][1] > maxVal {
			maxVal = arr[i][1]
		}
	}

	return -1, -1
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0]
	})

	fmt.Println(properties)

	startIndex, maxVal := findFistIndexDiff(properties)
	fmt.Println("DB ==>", startIndex, maxVal)

	maxNext := maxVal
	result := 0
	for i := startIndex; i >= 0; i-- {
		if properties[i][0] == properties[i+1][0] {
			if properties[i][1] < maxVal {
				fmt.Println("== ===>", i, properties[i][1], maxVal)
				result++
			}

			if properties[i][1] > maxNext {
				maxNext = properties[i][1]
			}

			continue
		}

		maxVal = maxNext
		maxNext = max(maxNext, properties[i][1])
		if properties[i][1] < maxVal {
			fmt.Println("!= ===>", i, properties[i][1], maxVal)
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(numberOfWeakCharacters([][]int{{1, 2}, {5, 10}, {5, 12}, {6, 13}}))
	fmt.Println("Hello world")
}
