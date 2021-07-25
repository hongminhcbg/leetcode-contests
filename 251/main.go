package main

import (
    "fmt"
)


func getIntFromS(s string) int {
    sum := 0
    for i := 0; i < len(s); i++ {
        sum += transform(int(s[i] - 'a' + 1))
    }

    return sum
}

func transform(in int) int {
    sum := 0
    for in != 0 {
        sum += in % 10
        in /= 10
    }

    return sum
}

func getLucky(s string, k int) int {
    valInt := getIntFromS(s)
    for i := 1; i < k; i++ {
        valInt = transform(valInt)
    }

    return valInt
}

func getUint8FromInt(in int) uint8 {
    s := fmt.Sprintf("%d", in)
    return s[0]
}


func maximumNumber(num string, change []int) string {
    arr := []uint8(num)
    canBreak := false
    for i := 0; i < len(arr); i++ {
        temp := getUint8FromInt(change[int(arr[i] - '0')])
        if arr[i] > temp && canBreak {
            break
        }

        if arr[i] < temp {
            arr[i] = temp
            canBreak = true
        }
    }

    return string(arr)
}

func check(s []int, m []int) int {
    sum := 0
    for i := 0; i < len(s); i++ {
        if s[i] == m[i] {
            sum++
        }
    }

    return sum
}

func calc(students [][]int, mentors [][]int, mentorOf []int) int {
    sum := 0
    for i := 0; i < len(mentorOf); i++ {
        sum += check(mentors[i], students[mentorOf[i]])
    }

    return sum
}

func next(arr []int) ([]int, bool) {
    n := len(arr)
    k := -1
    for i := n-2; i >=0; i-- {
        if arr[i] < arr[i+1] {
            k = i
            break
        }
    }

    if k == -1 {
        return nil, false
    }

    l := n-1
    for i := n-1; i > k; i-- {
        if arr[i] > arr[k] {
            l = i
            break
        }
    }

    arr[k], arr[l] = arr[l], arr[k]
    left := k + 1
    right := n - 1
    for {
        if !(left < right) {
            break
        }

        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }

    return arr, true
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    numP := len(students)
    mentorOf := make([]int, numP)
    for i := 0; i < numP; i++ {
        mentorOf[i] = i
    }

    max := -1
    isContinue := true
    for {
        fmt.Println(mentorOf)
        temp := calc(students, mentors, mentorOf)
        if temp > max {
            max = temp
        }

        mentorOf, isContinue = next(mentorOf)
        if !isContinue {
            break
        }
    }

    return max
}

func main(){
    fmt.Println("Hello world")
    //fmt.Println(getLucky("leetcode", 2))
    //fmt.Println(maximumNumber("5555888", []int{1,4,7,5,3,2,5,6,9,4}))
    mentors := [][]int{
        {1,0,1},
        {0,1,0},
        {1,1,1},
    }

    students := [][]int {
        {0,1,0},
        {1,0,1},
        {1,1,1},
    }

    fmt.Println(maxCompatibilitySum(mentors, students))
}
