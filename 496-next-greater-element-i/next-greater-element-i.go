func nextGreaterElement(nums1 []int, nums2 []int) []int {

    m := make(map[int]int)
    res := make([]int, len(nums2))
    st := []int{}
    finalres := []int{}

    // store index of nums2 elements
    for i, v := range nums2 {
        m[v] = i
    }
    res[len(nums2)-1] = -1
    st= append(st,nums2[len(nums2)-1])
    for i := len(nums2)-2 ; i>=0 ; i--{
        for (len(st)!=0 && nums2[i] >= st[len(st)-1]){
            st = st[:len(st)-1]
        }
        if (len(st)==0){
            res[i]=-1
        }else{
            res[i] = st[len(st)-1]
        }
        st = append(st, nums2[i])
    }

   for _, value := range nums1{
     i := m[value]
     v1 := res[i]
     finalres = append(finalres, v1)
   }
    return finalres
}