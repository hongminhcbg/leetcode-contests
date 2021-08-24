package main
import (
    "fmt"
)
func makeEqual(words []string) bool {
    chars := make(map[byte]int)
    for _, val := range words {
        for i := 0; i < len(val);  i++ {
            chars[val[i]] = chars[val[i]] + 1
        }
    }

    n := len(words)
    for _, val := range chars {
        if val % n != 0 {
            return false
        }
    }

    return true
}

func indexOf(a []byte, target byte) (int, bool) {
    for index, val := range a {
        if val == target {
            return index, true
        }
    }

    return -1, false
}

func findSub(s []byte, p []byte) ([]int, bool) {
    result := make([]int, len(p))
    index := 0
    for i := 0; i < len(p); i++ {
        indexX, ok := indexOf(s[index:], p[i])
        if !ok {
            return nil, false
        }

        result[i] = index + indexX
        index = index + indexX + 1
    }

    return result, true
}

func in(a []int, target int) bool {
    for _, val := range a {
        if val == target {
            return true
        }
    }

    return false
}

func maximumRemovals(s string, p string, removable []int) int {
    ss := []byte(s)
    pp := []byte(p)
    tracking, _ := findSub(ss, pp)
    fmt.Println("[DB] => tracking = ", tracking)
    var ok bool
    for i := 0; i < len(removable); i++ {
        ss[removable[i]] = '1'
        if !in(tracking, removable[i]) {
            continue
        }

        tracking, ok = findSub(ss, pp)
        fmt.Println("new tracking = ", tracking, ok, string(ss))
        if !ok {
            return i
        }
    }

    return len(removable)
}

func main(){
    fmt.Println("Hello world")
    //fmt.Println(makeEqual([]string{"abc", "abcc", "abd"}))
    fmt.Println(maximumRemovals("abcab", "abc", []int{0,1,2,3,4}))
}
