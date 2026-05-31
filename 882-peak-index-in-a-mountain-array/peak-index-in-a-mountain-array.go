func peakIndexInMountainArray(arr []int) int {
    // Search space is the entire array
    low := 0
    high := len(arr) - 1

    // Continue until only one element remains
    for low < high {
        mid := (low + high) / 2

        // If current element is smaller than the next,
        // we are on the ascending slope of the mountain.
        // Therefore, the peak must be to the right.
        if arr[mid] < arr[mid+1] {
            low = mid + 1
        } else {
            // We are on the descending slope or at the peak.
            // The peak could be mid itself, so keep mid
            // in the search space.
            high = mid
        }
    }

    // When low == high, both point to the peak index.
    return low
}