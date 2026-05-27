func searchRange(nums []int, target int) []int {

    // Initialize binary search pointers
    low := 0
    high := len(nums) - 1

    // Store first and last occurrence
    // Default = -1 if target not found
    ans1 := -1
    ans2 := -1

    // -----------------------------
    // Find FIRST occurrence
    // -----------------------------
    for low <= high {

        // Calculate middle index
        mid := low + (high-low)/2

        // Target found
        if nums[mid] == target {

            // Store possible first occurrence
            ans1 = mid

            // Continue searching on LEFT side
            // to find earlier occurrence
            high = mid - 1

        } else if target > nums[mid] {

            // Target is on right half
            low = mid + 1

        } else {

            // Target is on left half
            high = mid - 1
        }
    }

    // Reset pointers for second binary search
    low = 0
    high = len(nums) - 1

    // -----------------------------
    // Find LAST occurrence
    // -----------------------------
    for low <= high {

        // Calculate middle index
        mid := low + (high-low)/2

        // Target found
        if nums[mid] == target {

            // Store possible last occurrence
            ans2 = mid

            // Continue searching on RIGHT side
            // to find later occurrence
            low = mid + 1

        } else if target > nums[mid] {

            // Target is on right half
            low = mid + 1

        } else {

            // Target is on left half
            high = mid - 1
        }
    }

    // Return first and last positions
    return []int{ans1, ans2}
}