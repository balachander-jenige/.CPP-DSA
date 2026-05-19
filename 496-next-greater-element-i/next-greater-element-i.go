func nextGreaterElement(nums1 []int, nums2 []int) []int {

    m := make(map[int]int)
    res := []int{}

    // store index of nums2 elements
    for i, v := range nums2 {
        m[v] = i
    }

    for _, v := range nums1 {

        idx := m[v]
        found := -1

        // search to the right
        for i := idx + 1; i < len(nums2); i++ {
            if nums2[i] > v {
                found = nums2[i]
                break
            }
        }

        res = append(res, found)
    }

    return res
}