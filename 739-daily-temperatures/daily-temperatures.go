type Stack struct {
    data []int
}

// push element into stack
func (s *Stack) push(val int) {
    s.data = append(s.data, val)
}

// pop element from stack
func (s *Stack) pop() {
    s.data = s.data[:len(s.data)-1]
}

// get top element
func (s *Stack) top() int {
    return s.data[len(s.data)-1]
}

// check if stack is empty
func (s *Stack) empty() bool {
    return len(s.data) == 0
}
func dailyTemperatures(temp []int) []int {

    res := make([]int, len(temp))

    st := Stack{data: []int{}}

    for i := len(temp) - 1; i >= 0; i-- {

        // maintain monotonic decreasing stack
        for !st.empty() && temp[i] >= temp[st.top()] {
            st.pop()
        }

        if st.empty() {
            res[i] = 0
        } else {
            res[i] = st.top() - i
        }

        st.push(i)
    }

    return res
}