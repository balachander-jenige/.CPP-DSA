func subarraysDivByK(nums []int, k int) int {

      count :=0
    prefixsum :=0
    reminder := make(map[int]int)
    reminder[0] = 1

    for _, num := range nums {

        prefixsum += num
        rem := prefixsum % k 
        if rem <0 {
            rem = rem +k
        }
        if freq, exists := reminder[rem]; exists{
            count += freq
        }

        reminder[rem] +=1
    }    

    return count
    
}