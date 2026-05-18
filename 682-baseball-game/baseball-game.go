type Stack struct{
    items []int
}

func (s *Stack) pop(){
    if len(s.items) >0 {
        topIndex := len(s.items)-1
        s.items = s.items[:topIndex]
    }
}
func (s *Stack) peak() int{
    if len(s.items)>0{
          return s.items[len(s.items)-1]
    }
    return 0
}

func (s *Stack) push(v int){
    s.items= append(s.items,v)

}
func (s *Stack) Sum(){
   n := len(s.items) 
   total := s.items[n-1] + s.items[n-2]
   s.push(total)
}

func calPoints(operations []string) int {
    var s Stack

    for _,op := range operations{
        switch op {
            case "C" :
                s.pop()
            case "D" :
                s.push(s.peak()*2)
            case "+" :
                s.Sum()
            default:
                value, _ := strconv.Atoi(op)
                s.push(value)
                 
        }
    }
total :=0
    for _, v:= range s.items{
        total += v
    }

return total

}