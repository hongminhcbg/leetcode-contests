func isContain(a []int, target int) bool {
    for _, num := range a {
        if num == target {
            return true
        }
    }
    
    return false
}


func findDifference(nums1 []int, nums2 []int) [][]int {
    num1NotInNum2 := make([]int, 0)
    num2NotInNum1 := make([]int, 0)
    
    for _, n := range nums1 {
        if !isContain(nums2, n) && !isContain(num1NotInNum2, n) {
            num1NotInNum2 = append(num1NotInNum2, n)
        }
    }

    for _, n := range nums2 {
        if !isContain(nums1, n) && !isContain(num2NotInNum1, n) {
            num2NotInNum1 = append(num2NotInNum1, n)
        }
    }
    
    fmt.Println(num1NotInNum2, num2NotInNum1)
    return [][]int{num1NotInNum2, num2NotInNum1}
}

func choseNextIndex(nums []int, n int, currentIndex int) (int, int) {
    if currentIndex >= n {
        return -1, 0
    }
    
    nextIndex := currentIndex + 1
    deletedNum := 0
    for {
        if nextIndex >= n {
            break
        }
        
        if nums[nextIndex] != nums[currentIndex] {
            return nextIndex, deletedNum
        }
        
        nextIndex++
        deletedNum++
    }
    
    return -1, 0
}
func isAllElementEqual(a []int) bool {
    for i := 0; i < len(a) - 1; i++ {
        if a[i] != a[i+1] {
            return false
        }
    }
    
    return true
}

func minDeletion(nums []int) int {
    if isAllElementEqual(nums) {
        return len(nums)
    }
    
    result := 0 
    currentIndex := 0
    length := len(nums)
    
    for {
        nextIndex, deletedNum := choseNextIndex(nums, length, currentIndex)
        //fmt.Println(",current_index = ", currentIndex, ",next_index=", nextIndex, ",deletedNum=", deletedNum)
        if nextIndex == -1 {
            result += length - currentIndex
            break
        }
        
        result += deletedNum
        currentIndex = nextIndex + 1
    }
    
    return result
}

// return 10^n
var X = []int{1, 1e1, 1e2,1e3,1e4,1e5,1e6,1e7,1e8,1e9,1e10,1e11,1e12,1e13,1e14,1e15,1e16,1e17}
func mu(n int) int {
    return X[n]
    ans := 1
    for i := 0; i < n; i++ {
        ans *= 10
    }
    
    return ans
}

func dao(n int) int {
    ans := 0
    for {
        if n == 0 {
            break
        }
        
        tmp := n % 10
        ans = 10 * ans + tmp
        n /= 10
    }
    
    return ans
}

func calcMaxPalindrome(length int) int {
    if length % 2 == 0 {
        return mu(length/2) - mu(length/2-1) 
    }
    
    if length == 1 {
        return 9
    }
    
    length -= 1
    return (mu(length/2) - mu(length/2-1)) * 10
}

func kthPalindrome(queries []int, intLength int) []int64 {
    maxLength := calcMaxPalindrome(intLength)
    result := make([]int64, 0, len(queries))
    base := 0
    if intLength % 2 == 0 {
        base = mu(intLength/2-1)
    } else if intLength != 1 {
        base = mu(intLength/2-1)
    }
    
    //fmt.Println("maxLength = ", maxLength, ",base = ", base)

    
    for _, q := range queries {
        if q > maxLength {
            result = append(result, -1)
            continue
        }
        
        if intLength % 2 == 0 {
            ans := (base + q -1) * mu(intLength/2) + dao(base + q -1)
            result = append(result, int64(ans))
            continue
        }

        // intLength odd
        if intLength != 1 {
            mid := (q-1) % 10
            prefix := base + (q-1) / 10
            ans := prefix * mu(intLength/2+1) + mid * mu(intLength/2) + dao(prefix)
            result = append(result, int64(ans))
            continue
        }
        
        result = append(result, int64(q))   
    }
    
    return result
}
