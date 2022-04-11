func countCollisions(directions string) int {
    if len(directions) == 1 {
        return 0
    }
    
    left := 0
    right := len(directions) - 1
    for i := 0; i < len(directions); i++ {
        left = i 
        if directions[i] != 'L' {
            break
        }
    }
        
    for i := len(directions) - 1; i >= 0; i-- {
        right=i
        if directions[i] != 'R' {
            break
        }
    }
    
    if (left == len(directions) - 1 && directions[left] == 'L') || (right == 0 && directions[0] == 'R') {
        return 0
    }

    
    ans := 0
    for i := left; i <= right; i++ {
        if directions[i] != 'S' {
            ans += 1
        }
    }
    
    return ans
}

func findNextIndexDiffI(a []int, i int) int {
    for j := i+1; j < len(a); j++ {
        if a[j] != a[i] {
            return j
        }
    }
    
    return -1
}

func countHillValley(nums []int) int {
    ans := 0
    startIndex := 1 
    for {
        if startIndex >= len(nums) {
            break
        }
        
        rightIndex := findNextIndexDiffI(nums, startIndex)
        if rightIndex == -1 {
            break
        }
        
        if nums[startIndex-1] > nums[startIndex] && nums[rightIndex] > nums[startIndex] {
            // valley
            //fmt.Println("valley", startIndex, rightIndex-startIndex)
            ans += 1
            startIndex = rightIndex
            continue
        }
        
        if nums[startIndex-1] < nums[startIndex] && nums[rightIndex] < nums[startIndex] {
            // hill
            //fmt.Println("hill", startIndex, rightIndex-startIndex)
            ans += 1
            startIndex = rightIndex
            continue
        }

        startIndex = rightIndex
    }
    
    return ans
}
