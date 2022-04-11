
func largestInteger(num int) int {
    s := fmt.Sprintf("%d", num)
    b := []byte(s)
    for i := 0; i < len(b); i++ {
        maxIndex := i
        for j := i+1; j < len(b); j++ {
            if (b[j]-'0') %2 == (b[i]-'0')%2 && b[j] > b[maxIndex] {
                maxIndex = j
            }
        }
        
        if maxIndex != i {
            b[i], b[maxIndex] = b[maxIndex], b[i]
        }
    }
    
    ans, _ := strconv.Atoi(string(b))
    return ans
}

func multiStringInt(s string, i int) int {
    if len(s) == 0 {
        return i
    }
    
    sInt, _ := strconv.Atoi(s)
    return sInt*i
}

func addStr(a, b string) int {
        aInt, _ := strconv.Atoi(a)
        bInt, _ := strconv.Atoi(b)
    return aInt+bInt
}

// A*(B+C)*D
func minimizeResult(expression string) string {
    ans := "(" + expression + ")"
    args := strings.Split(expression, "+")
    min := addStr(args[0], args[1])
    
    left := args[0]
    for i := 0; i < len(left); i++ {
        for j := len(left)+1; j < len(expression); j++ {
            A := expression[0:i]
            D := expression[j+1:]
            args2 := strings.Split(expression[i:j+1], "+")
            B := args2[0]
            C := args2[1]
            
            currentAns := multiStringInt(D, multiStringInt(A, addStr(B,C)))
                        //fmt.Println(i,j,A,B,C,D, currentAns)
            if currentAns < min {
                min = currentAns
                ans = fmt.Sprintf("%s(%s+%s)%s", A, B, C, D)
            }
        }
    }
    
    return ans
}

func multiArrByMod(nums []int) int {
    const MOD = 1e9+7
    ans := 1
    for i := 0; i < len(nums); i++ {
        ans = (ans * nums[i]) % MOD
    }
    
    return ans
}

func maximumProduct(nums []int, k int) int {
    sort.Ints(nums)
    for {
        if k == 0 {
            break
        }
        
        nums[0] += 1
        k-- 
        for i := 1; i < len(nums); i++ {
            if k == 0 {
                break
            }
            
            if nums[i] < nums[i-1] {
                nums[i] += 1
                k--
                continue
            }
            break
        }
    }
    
    //fmt.Println(nums)
    return multiArrByMod(nums)
}
