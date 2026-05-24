type Stack struct {
    data []int // underlying slice that stores stack elements
}

// push element into stack
func (s *Stack) push(val int) {
    // append value at the end (top of stack)
    s.data = append(s.data, val)
}

// pop element from stack
func (s *Stack) pop() {
    // remove last element (top of stack)
    s.data = s.data[:len(s.data)-1]
}

// get top element
func (s *Stack) top() int {
    // return last element without removing it
    return s.data[len(s.data)-1]
}

// check if stack is empty
func (s *Stack) empty() bool {
    // stack is empty when slice length is 0
    return len(s.data) == 0
}

func nextGreaterElements(nums []int) []int {

    var st Stack // stack to maintain next greater candidates
    res := make([]int, len(nums)) // result array

    // preload stack with elements from index 0 to n-2
    // (attempt to simulate circular behavior)
    for i := len(nums)-2; i >= 0; i-- {
        st.push(nums[i])
    }

    // traverse array from right to left
    for i := len(nums)-1; i >= 0; i-- {

        // remove all elements smaller or equal to current element
        // because they cannot be next greater
        for !st.empty() && nums[i] >= st.top() {
            st.pop()
        }

        // if stack is empty, no greater element exists
        if st.empty() {
            res[i] = -1
        } else {
            // top of stack is next greater element
            res[i] = st.top()
        }

        // push current element into stack for future comparisons
        st.push(nums[i])
    }

    return res
}