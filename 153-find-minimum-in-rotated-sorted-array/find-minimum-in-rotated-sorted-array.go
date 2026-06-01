func findMin(nums []int) int {
    n := len(nums)

    // Binary search range
    low := 0
    high := n - 1

    // Stores index of potential minimum element
    res := 0

    for low <= high {
        mid := (high + low) / 2

        // If middle element is greater than the last element,
        // it means the minimum is in the right half (rotated part)
        if nums[mid] > nums[n-1] {
            low = mid + 1
        } else {
            // nums[mid] could be the minimum, so store it
            res = mid

            // Try to find a smaller value on the left side
            high = mid - 1
        }
    }

    // Return the minimum value found
    return nums[res]
}