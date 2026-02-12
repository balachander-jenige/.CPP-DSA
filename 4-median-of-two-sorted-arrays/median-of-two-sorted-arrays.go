func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1 :=len(nums1)
    l2 :=len(nums2)

    var i,j int
    var merged = []int{}
    
    for i < l1 && j < l2 {
        if nums1[i] < nums2[j] {
            merged = append(merged, nums1[i])
            i++
        } else {
            merged = append(merged, nums2[j])
            j++
        }
    }

    // append remaining elements
    for i < l1 {
        merged = append(merged, nums1[i])
        i++
    }

    for j < l2 {
        merged = append(merged, nums2[j])
        j++
    }

    index := len(merged)

    if index%2 !=0 {
        return float64(merged[index/2])
    }else{
        index:= index/2
        return (float64(merged[index-1]+merged[index])/2)
    }


 }






    
