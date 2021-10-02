package main

import (
	"fmt"
	"strconv"
	"strings"
)

func finalValueAfterOperations(operations []string) int {
	op := map[string]int{
		"++X": 1,
		"X++": 1,
		"X--": -1,
		"--X": -1,
	}

	result := 0
	for _, o := range operations {
		result += op[o]
	}

	return result
}

func sumOfBeauties(nums []int) int {
	maxFromLeft := make([]int, len(nums))
	minFromRight := make([]int, len(nums))

	n := len(nums)
	maxFromLeft[0] = nums[0]
	// init
	for i := 1; i < n; i++ {
		if maxFromLeft[i-1] < nums[i] {
			maxFromLeft[i] = nums[i]
		} else {
			maxFromLeft[i] = maxFromLeft[i-1]
		}
	}

	minFromRight[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] < minFromRight[i+1] {
			minFromRight[i] = nums[i]
		} else {
			minFromRight[i] = minFromRight[i+1]
		}
	}

	result := 0
	for i := 1; i < len(nums)-1; i++ {
		if maxFromLeft[i-1] < nums[i] && nums[i] < minFromRight[i+1] {
			result += 2
			continue
		}

		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			result += 1
			continue
		}
	}

	return result
}

func xy2key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func key2xy(s string) (int, int) {
	args := strings.Split(s, ":")
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])

	return x, y
}

type DetectSquares struct {
	points map[string]int
}

func Constructor() DetectSquares {
	points := make(map[string]int)
	return DetectSquares{
		points: points,
	}
}

func (this *DetectSquares) Add(point []int) {
	key := xy2key(point[0], point[1])
	val, ok := this.points[key]
	if ok {
		this.points[key] = val + 1
	} else {
		this.points[key] = 1
	}
}

func (this *DetectSquares) Count(point []int) int {
	x := point[0]
	y := point[1]
	result := 0

	for key, _ := range this.points {
		x1, y1 := key2xy(key)
		if x1 == x || y1 == y {
			continue
		}

		countx1y1 := this.points[key]
		keyxy1 := xy2key(x, y1)
		countxy1, ok := this.points[keyxy1]
		if !ok {
			continue
		}

		keyx1y := xy2key(x1, y)
		countx1y, ok := this.points[keyx1y]
		if !ok {
			continue
		}

		fmt.Println(xy2key(x, y), key, countx1y1, countx1y, countxy1)
		result += countx1y1 * countx1y * countxy1

	}

	return result
}

func main() {
	fmt.Println("Hello world")
	//fmt.Println(finalValueAfterOperations([]string{"++X", "X++", "--X", "X--"}))
	//fmt.Println(sumOfBeauties([]int{1, 2, 3, 4, 5, 56, 67, 67}))
	obj := Constructor()
	obj.Add([]int{5, 10})
	obj.Add([]int{10, 5})
	obj.Add([]int{10, 10})
	obj.Add([]int{3, 0})
	obj.Add([]int{8, 0})
	obj.Add([]int{8, 5})

	fmt.Println(obj.Count([]int{3, 5}))
	obj.Add([]int{9, 0})
	obj.Add([]int{9, 8})
	obj.Add([]int{1, 8})

	obj.Add([]int{0, 0})
	obj.Add([]int{8, 0})
	obj.Add([]int{8, 8})

	fmt.Println(obj.Count([]int{0, 8}))
	//	obj.Add([]int{11, 2})
	fmt.Println(obj.Count([]int{11, 10}))

}
