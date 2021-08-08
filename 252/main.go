package main
import (
    "fmt"
    "sort"
)

func isThree(n int) bool {
    sum := 2
    for i := 2; i < n; i++ {
        if sum > 3 {
            return false
        }

        if n % i == 0 {
            sum++
        }
    }

    return sum == 3
}

func numberOfWeeks(milestones []int) int64 {
    var sum int64 = 0
    sort.Ints(milestones)

    offset := 0
    n := len(milestones)
    for i := 0; i < len(milestones); i++ {
        if milestones[i] == offset {
            continue
        }

        if i == n - 1 {
            // LAST index
            delta := int64(milestones[i] - offset)
            max := sum - 2*int64(offset) + 1
            if delta < max {
                sum = sum + delta
                continue
            }

            sum = sum + max
            continue
        }

        sum += int64(n-i) * int64(milestones[i]-offset)
        offset = milestones[i]
    }

    return sum
}

const MAX_APPLE = 1e15

func calcAppleInlayer(i int64) int64{
    return (3*i) * (i+1) * 4 - 12*i
}

func calcEdgeLen(apples []int64, neededApple int64) ( bool, int) {
    left := 0
    right := len(apples) - 1
    for {
        fmt.Println(left, right)
        if right - left == 1 {
           return false, right
        }

        mid := (left + right) / 2
        if apples[mid] == neededApple {
            return true, mid
        }

        if apples[mid] > neededApple {
            right = mid
            continue
        }

        left = mid
    }
}

func minimumPerimeter(neededApples int64) int64 {
    apples := make([]int64, 1)
    apples[0] = 0
    layer := 1
    var appleInlayer int64 = 0
    for {
        appleInlayer = calcAppleInlayer(int64(layer))
        apples = append(apples, apples[layer-1] + appleInlayer)
        if appleInlayer > neededApples {
            break
        }

        fmt.Println("[DB] ===>", appleInlayer)

        layer++
    }
    
    fmt.Println("[DB] ===>", apples) 
    _, choseLayer := calcEdgeLen(apples, neededApples)
    return int64(choseLayer) * 8
}

func main(){
    fmt.Println("Hello world")
    //fmt.Println(isThree(5))
    //fmt.Println(numberOfWeeks([]int{5,1,2}))
    fmt.Println(minimumPerimeter(1000))
}
