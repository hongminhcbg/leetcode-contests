package main

import "fmt"

func calcSumDigit(n int) int {
	sum := 0
	for {
		if n == 0 {
			break
		}

		sum += n % 10
		n /= 10
	}
	return sum
}

func countEven(num int) int {
	result := 0
	for i := 2; i <= num; i++ {
		if calcSumDigit(i) % 2 == 0 {
			result += 1
		}
	}
	return result
}

type ListNode struct {
	     Val int
	     Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}

	var tail *ListNode
	tail = result
	sum := 0

	head = head.Next
	for {
		if head == nil {
			break
		}

		if head.Val != 0 {
			sum += head.Val
			head = head.Next
			continue
		}

		if result.Val == 0 {
			result = &ListNode{
				Val:  sum,
				Next: nil,
			}

			tail = result
			head = head.Next
			sum = 0
			continue
		}

		newTail := &ListNode{
			Val:  sum,
			Next: nil,
		}

		tail.Next = newTail
		tail = newTail
		sum = 0
		head = head.Next
	}

	return result
}

func findBiggestCharDiffLastChar(chars []int, lastChard uint8) (uint8, uint8) {
	max := uint8(26)
	selected := uint8(100)

	for i := 25; i >= 0; i-- {
		if chars[i] > 0 && max == 26 {
			max = uint8(i)
		}

		if chars[i] > 0 && i != int(lastChard) {
			selected = uint8(i)
			return selected, max
		}
	}

	return uint8(100), max
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func appendChar(s []uint8, c uint8, num int) []uint8 {
	temp := make([]uint8, 0, num)
	for i := 0; i < num; i++ {
		temp = append(temp, c+'a')
	}

	return append(s, temp...)
}

func repeatLimitedString(s string, repeatLimit int) string {
	result := make([]uint8, 0)
	chars := make([]int, 26)
	for _, c := range s {
		chars[int(c-'a')]++
	}

	lastChar := uint8(26)
	var num int
	for {
		selectedChar, maxChar := findBiggestCharDiffLastChar(chars, lastChar)
		if selectedChar == 100 {
			break
		}

		if selectedChar == maxChar {
			num = min(chars[selectedChar], repeatLimit)
		} else {
			num = 1
		}

		chars[selectedChar] -= num
		fmt.Println("selected char = ", selectedChar, ", maxchar = ", maxChar, ", num = ", num)
		lastChar = selectedChar
		result = appendChar(result, selectedChar, num)
	}

	return string(result)
}

func ucln(a, b int) int {
	for {
		if a % b == 0 {
			return b
		}

		if b % a == 0 {
			return a
		}

		if a > b {
			a = a % b
			continue
		}

		b = b % a
	}
}

func countPairs(nums []int, k int) int64 {
	hash := make(map[int]int64)
	n := int64(len(nums))
	for _, num := range nums {
		hash[num%k] += 1
	}
	//fmt.Printf("hash = %+v\n", hash)
	result := hash[0] * (n-hash[0])
	//  fmt.Println("result1 = ", result)

	if hash[0] >= 2 {
		result += (hash[0] * (hash[0]-1))/2
	}
	//  fmt.Println("result2 = ", result)


	var sub int
	for i := 2; i < k; i++ {
		if hash[i] == 0 {
			continue
		}

		ucln := ucln(k, i)
		if ucln == 1 {
			continue
		}

		sub = k / ucln
		subInterator := sub

		for {
			if subInterator >= k {
				break
			}

			if subInterator < i {
				subInterator += sub
				continue
			}
			if subInterator == i {
				if hash[subInterator] > 0 {
					result += (hash[subInterator] * (hash[subInterator] - 1))/2
				}
				subInterator += sub
				continue
			}

			result += hash[i] * hash[subInterator]
			subInterator += sub
		}
	}
	//fmt.Println("result3 = ", result)

	return result
}


func main(){
	//fmt.Println(countEven(30))
	fmt.Println(repeatLimitedString("cccccccbbbbaa", 2))
}
