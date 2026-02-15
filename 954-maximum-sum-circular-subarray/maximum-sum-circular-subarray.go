func maxSubarraySumCircular(nums []int) int {

    globalmax :=nums[0]
    globalmin :=nums[0]

    currmax :=0
    currmin :=0
    totalsum :=0

    for _,n := range nums{
        currmax = max(currmax + n, n)
        currmin = min(currmin + n, n)

        globalmax = max(currmax, globalmax)  //linear max
        globalmin = min(currmin, globalmin)  // linear min

        totalsum +=n

    }

        if globalmax > 0{
            return max(globalmax, totalsum - globalmin)  //toget circularmax = totalsum - globalmin
        }

        return globalmax

     
}