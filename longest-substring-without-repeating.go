func lengthOfLongestSubstring(s string) int {
    n := len(s)
    res := 0 
    chMap := make(map[string]int)
    i :=0
    for j := 0 ; j < n; j++{
        if val, ok := chMap[string(s[j])]; ok {
            i = max(val, i)
        }
        res  = max(res, j-i +1)
        chMap[string(s[j])] = j +1
    }
    return res 
}

func max(x int, y int)int {
    if x > y {
        return x
    }
    return y
} 