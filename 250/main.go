package main
import (
    "fmt"
    "strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
    result := 0
    args := strings.Split(text, " ")

    for _, str := range args {
        ok := true
        for i := 0; i < len(brokenLetters); i++ {
            if strings.Contains(str, brokenLetters[i:i+1]) {
                ok = false
                break
            }
        }

        if ok {
            result++
        }
    }

    return result
}

func caclRung(src, dest, dist int) int {
    delta := dest - src
    if delta <= dist {
        return 0
    }

    if delta % dist == 0 {
        return delta / dist - 1
    }

    return delta / dist
}

func addRungs(rungs []int, dist int) int {
    rungs = append([]int{0}, rungs...)
    result := 0
    for i := 0; i < len(rungs) - 1; i++ {
        result += caclRung(rungs[i], rungs[i+1], dist)
    }

    return result
}

func abs(x, y int) int {
    if x > y {
        return x - y
    }

    return y - x
}

func maxPoints(points [][]int) int64 {
    m := len(points)
    n := len(points[0])
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            max := 0
            for index := 0; index < n; index++ {
                if (points[i][j] + points[i-1][index] - abs(j, index) > max) {
                    max = points[i][j] + points[i-1][index] - abs(j, index)
                }
            }

            points[i][j] = max
        }
    }

    fmt.Println("[DB] ======> new points = ", points)

    result := 0
    for i := 0; i < n; i++{
        if points[m-1][i] > result {
            result = points[m-1][i]
        }
    }

    return int64(result)
}

func main(){
    fmt.Println("Hello world")
//    fmt.Println(canBeTypedWords("Leet code", "d"))
//    rungs := []int{1,3,5,10}
//    fmt.Println(addRungs(rungs, 2))
    points := [][]int{{1,2,3}, {1,5,1}, {3,1,1}}
    fmt.Println(maxPoints(points))
}
