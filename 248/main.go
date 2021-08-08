package main 
import (
    "fmt"
    "sort"
)

func maxProductDifference(nums []int) int {
    sort.Ints(nums)
    lenngth := len(nums)
    return nums[lenngth-1] * nums[lenngth-2] - nums[0] * nums[1]
}

func isWonderful(in map[uint8] int) bool {
    sumOdd := 0
    for _, val := range in {
        //fmt.Println(key, val)
        if val % 2 == 1 {
            sumOdd++
        }
    }

    return sumOdd <= 1
}

func wonderfulSubstrings(word string) int64 {
    trackingWord := make(map[uint8] int)
    var i uint8
    for i = 0; i < 10; i++ {
        trackingWord[i] = 0
    }
    
    result := int64(0)
    for i := 0; i < len(word); i++ {
        for key, _ := range trackingWord {
            trackingWord[key] = 0
        }

        for j := i; j < len(word); j++ {
            trackingWord[word[j] - 'a'] = trackingWord[word[j] - 'a'] + 1
            if isWonderful(trackingWord) {
                result++
            }

            //fmt.Println("[DB] =====> ", i, j, word[i:j+1])
        }
    }

    return result
}

func main(){
    fmt.Println("Hello world")
    fmt.Println(wonderfulSubstrings("aba"))
}
