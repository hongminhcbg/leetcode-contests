package main

import (
	"fmt"
)

func reverseStr(s string) string {
	b := make([]byte, len(s), len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b[len(s)-1-i] = s[i]
	}

	return string(b)
}

func reversePrefix(word string, ch byte) string {
	index := -1
	b := []byte(word)
	for i := 0; i < len(b); i++ {
		if b[i] == ch {
			index = i
			break
		}
	}

	if index == -1 {
		return word
	}

	return reverseStr(word[:index+1]) + word[index+1:]
}

func ucln(a, b int) int {
	for {
		if a%b == 0 {
			return b
		}

		if b%a == 0 {
			return a
		}

		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
}

func toPsString(a, b int) string {
	uc := ucln(a, b)
	return fmt.Sprintf("%d/%d", a/uc, b/uc)
}

func interchangeableRectangles(rectangles [][]int) int64 {
	n := len(rectangles)
	countRatio := make(map[string]int64)
	for i := 0; i < n; i++ {
		str := toPsString(rectangles[i][0], rectangles[i][1])
		countRatio[str] = countRatio[str] + 1
	}

	fmt.Printf("[DB] ===> %+v ", countRatio)
	result := int64(0)
	for _, val := range countRatio {
		result += val * (val - 1) / 2
	}

	return result
}

func main() {
	fmt.Println("Hello world")
	//fmt.Println(reversePrefix("dsfdfdffd", 'f'))
	fmt.Println(interchangeableRectangles([][]int{{2, 4}, {3, 6}, {100, 200}, {7, 3}}))
}
