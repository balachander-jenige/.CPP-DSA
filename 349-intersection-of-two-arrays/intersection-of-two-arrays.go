func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1)==0||len(nums2)==0{
        return []int{}
    }
    var res = []int{}
    var mapa = make(map[int]bool)
    for _,v:= range nums1{
        mapa[v]=true
    }
    for _,v:= range nums2{
        if mapa[v]{
            res= append(res,v)
            mapa[v]=false
        }
    }
    return res
}