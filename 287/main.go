package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertHHMM2int(s string) int {
	args := strings.Split(s, ":")
	h, _ := strconv.Atoi(args[0])
	m, _ := strconv.Atoi(args[1])
	return 60*h + m
}

func convertTime(current string, correct string) int {
	currentInt := convertHHMM2int(current)
	correctInt := convertHHMM2int(correct)
	result := 0
	sub := correctInt - currentInt
	for {
		if sub == 0 {
			break
		}

		if sub >= 60 {
			sub -= 60
			result++
			continue
		}

		if sub >= 15 {
			sub -= 15
			result++
			continue
		}

		if sub >= 5 {
			sub -= 5
			result++
			continue
		}

		sub -= 1
		result++
		fmt.Println(sub, result)
	}

	return result
}

func findWinners(matches [][]int) [][]int {
	const MAX = 1e5
	n := len(matches)
	cntLose := make([]int, MAX+1)
	cntWin := make([]int, MAX+1)
	for i := 0; i < n; i++ {
		winI, loseI := matches[i][0], matches[i][1]
		cntLose[loseI]++
		cntWin[winI]++
	}

	//fmt.Println("win list = ", cntWin)
	//fmt.Println("Lose list = ", loseI)
	onlyWinList := make([]int, 0)
	loseOnlyOneList := make([]int, 0)
	for i := 0; i < MAX+1; i++ {
		if cntWin[i] > 0 && cntLose[i] == 0 {
			onlyWinList = append(onlyWinList, i)
			continue
		}

		if cntLose[i] == 1 {
			loseOnlyOneList = append(loseOnlyOneList, i)
		}
	}

	return [][]int{onlyWinList, loseOnlyOneList}
}

func maximumCandies(candies []int, k int64) int {
	sum := 0
	for _, val := range candies {
		sum += val
	}

	if sum < int(k) {
		return 0
	}

	l := 1
	r := sum / int(k)
	ans := 0
	mid := 0
	cntPipe := 0
	for {
		if l > r {
			break
		}

		cntPipe = 0
		mid = l + (r-l)/2
		for i := 0; i < len(candies); i++ {
			cntPipe += candies[i] / mid
		}

		if cntPipe >= int(k) {
			ans = mid
			l = mid + 1
			continue
		}

		r = mid - 1
	}

	return ans
}

func main() {
	//fmt.Println("Hello world")
	//fmt.Println(convertTime("01:15", "02:15"))
	fmt.Println(maximumCandies([]int{4, 7, 5}, 16))
}
