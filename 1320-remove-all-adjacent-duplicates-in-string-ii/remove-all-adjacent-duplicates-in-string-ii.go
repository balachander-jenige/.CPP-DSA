type Pair struct {
    ch    rune
    count int
}

type Stack struct {
    data []Pair
}

// push element into stack
func (s *Stack) push(val Pair) {
    s.data = append(s.data, val)
}

// pop element from stack
func (s *Stack) pop() {
    s.data = s.data[:len(s.data)-1]
}

// get top element
func (s *Stack) top() Pair {
    return s.data[len(s.data)-1]
}

// check if stack is empty
func (s *Stack) empty() bool {
    return len(s.data) == 0
}
func removeDuplicates(s string, k int) string {

    var st Stack

    // traverse each character
    for _, ch := range s {

        // if stack not empty and same character
        if !st.empty() && st.top().ch == ch {

            // increase frequency
            st.data[len(st.data)-1].count++

            // remove if count becomes k
            if st.top().count == k {
                st.pop()
            }

        } else {

            // push new character with count 1
            st.push(Pair{
                ch:    ch,
                count: 1,
            })
        }
    }

    // build final string
    res := ""

    for _, p := range st.data {

        for i := 0; i < p.count; i++ {
            res += string(p.ch)
        }
    }

    return res
}