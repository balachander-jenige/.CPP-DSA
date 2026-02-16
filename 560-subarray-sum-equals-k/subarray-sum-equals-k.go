func subarraySum(nums []int, k int) int {
    
    count :=0
    prefixsum :=0
    prefixcount := make(map[int]int)
    prefixcount[0] = 1

    for _, num := range nums {

        prefixsum += num

        if freq, exists := prefixcount[prefixsum - k]; exists{
            count += freq
        }

        prefixcount[prefixsum] +=1
    }    

    return count
}