func divisorSubstrings(num int, k int) int {
    ans := 0
    numStr := fmt.Sprintf("%d", num)
    for i := 0; i <= len(numStr) - k; i++ {
        subStr := numStr[i: i+k]
        tmp, _ := strconv.Atoi(subStr)
        
        if tmp == 0 {
            continue
        }
        
        if num % tmp == 0 {
            ans += 1
        }
    }
    
    return ans
}
func calcTotal(nums []int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    
    return s
}

func waysToSplitArray(nums []int) int {
    sumAll := calcTotal(nums)
    ans := 0
    currentSumLeft := 0
    currentSumRight := 0
    for i := 0; i < len(nums) - 1; i++ {
        currentSumLeft += nums[i]
        currentSumRight = sumAll - currentSumLeft
        if currentSumLeft >= currentSumRight {
            ans += 1
        }
    }
    
    return ans
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    
    return b
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
    n := len(tiles)
    sort.SliceStable(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })
    left := tiles[0][0]
    right := left + carpetLen - 1
    right = min(right, tiles[n-1][1])
    //fmt.Println(tiles)
    currentTilesIndexLeft := 0
    currentTilesIndexRight := 0
    currentCnt := 0
    for i := left; i <= right; i++ {
        if tiles[currentTilesIndexRight][0] <= i && i <= tiles[currentTilesIndexRight][1] {
           currentCnt += 1
            continue
        }
        
        // outside currentTilesIndexRight
        if currentTilesIndexRight == n - 1 {
            continue
        }
        
        if tiles[currentTilesIndexRight+1][0] <= i && i <= tiles[currentTilesIndexRight+1][1] {
           currentCnt += 1
            currentTilesIndexRight += 1
            continue
        }
    }
    
    ans := currentCnt
    //fmt.Println("init finished", left, right, currentTilesIndexLeft, currentTilesIndexRight, ans)
    
    // loop right -> max tiles
    
    for {
        right += 1
        if right > tiles[n-1][1] {
            break
        }
        
        if tiles[currentTilesIndexRight][0] <= right && right <= tiles[currentTilesIndexRight][1] {
            currentCnt += 1
        }
        if currentTilesIndexRight < n -1 {
            if tiles[currentTilesIndexRight+1][0] <= right && right <= tiles[currentTilesIndexRight+1][1] {
                currentCnt += 1
                currentTilesIndexRight += 1
            }
        }
        
        if tiles[currentTilesIndexLeft][0] <= left && left <= tiles[currentTilesIndexLeft][1] {
            currentCnt -= 1
        }
        if currentTilesIndexLeft < n - 1 {
            if tiles[currentTilesIndexLeft+1][0] <= left && left <= tiles[currentTilesIndexLeft+1][1] {
                currentCnt -= 1
                currentTilesIndexLeft += 1
            }
        }
        left += 1
        //fmt.Println("===>", left, right, currentCnt, currentTilesIndexLeft, currentTilesIndexRight)
        if currentCnt > ans {
            ans = currentCnt
        }
    }
    
    return ans
    
}
