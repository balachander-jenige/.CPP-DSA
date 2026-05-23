func dailyTemperatures(temp []int) []int {
    res := make([]int, len(temp))
    m := make(map[int]int)
    st := []int{}
    
    for i:= len(temp)-1 ;i>=0; i--{
        for len(st) >0  && temp[i] >= st[len(st)-1] { 
            st = st[:len(st)-1]
        }
        if len(st) == 0{
            res[i] = 0

        }else{
            res[i] = m[st[len(st)-1]] -i
        }
        st =append(st,temp[i])
        m[temp[i]] = i
    }
    return res
    
} 