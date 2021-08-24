package main
import (
    "fmt"
)
func GCD(a int, b int) int {
    for a != b {
        if a % b == 0 {
            return b
        }

        if b % a == 0 {
            return a
        }

        if a > b {
            a = a % b
        } else {
            b = b % a
        }
    }

    return a
}

func findGCD(nums []int) int {
    max := nums[0]
    min := nums[0]
    for _, val := range nums {
        if max < val {
            max = val
        }

        if min > val {
            min = val
        }
    }
    
    return GCD(max, min)
}

func increment(s string) string {
    result := []byte(s)

    remember := byte(1)
    for i := len(result) - 1; i >=0; i-- {
        temp := result[i] + remember
        result[i] = (temp%2) + '0'
        remember = temp / 2
    }

    return string(result)
}

func contain(nums []string, target string) bool {
    for _, val := range nums {
        if val == target {
            return true
        }
    }

    return false
}

func findDifferentBinaryString(nums []string) string {
    start := ""
    n := len(nums)
    for i := 0; i < n; i++ {
        start = fmt.Sprintf("%s%d", start, 0)
    }

    for {
        fmt.Println(start)
        if !contain(nums, start) {
            return start
        }

        start = increment(start)
    }
}
func main(){
    fmt.Println("Hello world")
    //fmt.Println(findGCD([]int{4,6,8,12}))
    fmt.Println(findDifferentBinaryString([]string{"000", "001", "010"}))
}
