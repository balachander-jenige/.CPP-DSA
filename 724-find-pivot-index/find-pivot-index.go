func pivotIndex(nums []int) int {

    totalsum := 0

    prefixsum := 0

    for _, num := range nums {
        totalsum +=num
    }

    for idx, num := range nums {
        right := totalsum - prefixsum -num 

        if prefixsum == right {
            return idx
        }

        prefixsum += num
    }

    return -1;
    
}