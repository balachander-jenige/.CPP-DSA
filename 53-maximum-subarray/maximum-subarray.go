func maxSubArray(nums []int) int {
    
    bestending := nums[0]
    ans :=nums[0]

    for i :=1; i< len(nums); i++ {

        v1 := bestending + nums[i]
        v2 := nums[i]

        bestending = max(v1,v2)        
        ans = max(ans, bestending)

    }
    return ans
}