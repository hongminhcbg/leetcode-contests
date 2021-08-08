package main
import (
    "fmt"
)

func isPrefixString(s string, words []string) bool {
    sumStr := ""
    for _, val := range words {
        if len(sumStr) >= len(s) {
            break
        }
        sumStr = sumStr + val
    }

    return sumStr == s
}
func minSwaps(s string) int {
    numClose := 0
    numOpen := 0  
    for i := 0; i < len(s); i++ { 
        if s[i] == '[' { 
            numOpen++
            continue
        }
        if s[i] == ']' { 
            if numOpen > 0 { 
                numOpen-- 
                continue
            }
         
            numClose++
        }
    }
    if numClose % 2 == 0 { 
        return numClose / 2
    } 
    
    return numClose / 2 + 1
}

func main(){
    fmt.Println("Hello world")
    fmt.Println(minSwaps("][][]][[[[]]"))
}
