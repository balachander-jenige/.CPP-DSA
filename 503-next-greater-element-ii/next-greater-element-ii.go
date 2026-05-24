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
func nextGreaterElements(nums []int) []int {
    var st Stack
    res := make([]int,(len(nums)))
    for i:= len(nums)-2;i>=0;i--{
        st.push(nums[i])
    }
for i:= len(nums)-1; i>=0; i--{
    for !st.empty() && nums[i]>= st.top(){
        st.pop()
    }
    if st.empty(){
        res[i]= -1
    }else{
        res[i]=st.top()
    }
    st.push(nums[i])
}
return res
    
}
 


