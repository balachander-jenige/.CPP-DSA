func findMin(nums []int) int {
    n := len(nums)
    low := 0
    high := len(nums)-1
    res := 0

    for low <= high{
        mid := (high+low)/2
        if nums[mid] > nums[n-1]{
            low = mid+1
        }else{
            res = mid
            high = mid -1
        }
    }
    return nums[res]
}