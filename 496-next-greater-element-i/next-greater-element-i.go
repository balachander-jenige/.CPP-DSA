func nextGreaterElement(nums1 []int, nums2 []int) []int {

    // Map to store:
    // key   -> element from nums2
    // value -> next greater element
    nextGreater := make(map[int]int)

    // Monotonic decreasing stack
    // Stores elements from nums2
    stack := []int{}

    // Traverse nums2 from right to left
    for i := len(nums2) - 1; i >= 0; i-- {

        // Remove all smaller or equal elements
        // because they cannot be the next greater element
        for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {

            // Pop from stack
            stack = stack[:len(stack)-1]
        }

        // If stack is empty,
        // no greater element exists on right side
        if len(stack) == 0 {

            nextGreater[nums2[i]] = -1

        } else {

            // Top of stack is the next greater element
            nextGreater[nums2[i]] = stack[len(stack)-1]
        }

        // Push current element into stack
        stack = append(stack, nums2[i])
    }

    // Result array for nums1
    res := make([]int, len(nums1))

    // Fetch answers for nums1 elements
    for i, v := range nums1 {

        // Directly get next greater element from map
        res[i] = nextGreater[v]
    }

    // Return final answer
    return res
}