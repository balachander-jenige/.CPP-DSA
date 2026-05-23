func dailyTemperatures(temp []int) []int {

    // Result array to store number of days until a warmer temperature
    res := make([]int, len(temp))

    // Map to store last seen index of each temperature value
    m := make(map[int]int)

    // Stack to store temperatures (not indexes in this version)
    st := []int{}

    // Traverse the temperature array from right to left
    for i := len(temp) - 1; i >= 0; i-- {

        // Pop elements from stack while current temperature
        // is greater or equal to stack top temperature
        for len(st) > 0 && temp[i] >= st[len(st)-1] {
            st = st[:len(st)-1]
        }

        // If no warmer temperature exists to the right
        if len(st) == 0 {
            res[i] = 0
        } else {
            // Use map to get index of next warmer temperature
            // and compute difference in days
            res[i] = m[st[len(st)-1]] - i
        }

        // Push current temperature into stack
        st = append(st, temp[i])

        // Store index of current temperature
        m[temp[i]] = i
    }

    return res
}