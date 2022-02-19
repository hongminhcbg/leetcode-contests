package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for i, _ := range nums {
		if nums[i] == original {
			original *= 2
		}
	}

	return original
}

func sumArr(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n == 1 {
			sum +=1
		}
	}
	return sum
}

func maxScoreIndices(nums []int) []int {
	result := []int{0}
	max := sumArr(nums)
	previousScore := max
	currentScore := 0
	for i := 1; i <= len(nums); i++ {
		 if nums[i-1] == 0 {
			currentScore = previousScore+1
		} else {
			currentScore = previousScore-1
		}

		if currentScore > max {
			max = currentScore
			result = []int{i}
			previousScore = currentScore
			continue
		}

		if currentScore == max {
			result = append(result, i)
		}

		previousScore = currentScore
	}

	return result
}


var p []int
func getCharInt(c uint8) int {
	return int(c - 'a' + 1)
}

func calcOneString(s string, power int, k int, m int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = (result + getCharInt(s[i]) * p[i]) % m
	}

	return result
}

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	p = make([]int, k, k)
	p[0] = 1
	for i := 1; i < k; i++ {
		p[i] = (p[i-1] * power) % modulo 
	}
	
	for i := 0; i <= len(s) - k; i++ {
		hashI := calcOneString(s[i:i+k], power, k, modulo)
		if hashI == hashValue {
			return s[i:i+k]
		}
	}

	return ""
}

func main() {
	//fmt.Println(findFinalValue([]int{1,2,4,8,16}, 5))
	//fmt.Println("vim-go")
}
