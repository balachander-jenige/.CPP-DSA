func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    //Inserting at interval
    var result [][]int;
    inserterd := false
    for i:=0 ;i<len(intervals);i++{
        if (!inserterd && intervals[i][0] > newInterval[0]){
            result = append(result, newInterval)
            inserterd = true
        }
        result = append(result,intervals[i])
    }
    if !inserterd{
        result = append(result,newInterval)
    }

    intervals = result
    result = [][]int{} 
    start1 := intervals[0][0]
    end1   := intervals[0][1]    
   
    for i:= 1; i<len(intervals) ; i++ {
        start2 := intervals[i][0]
        end2   := intervals[i][1]
         if end1 >= start2 {
            start1 = start1
            end1   = max(end1, end2)
         }else{
         result = append(result, [] int{start1,end1})
         start1 = start2
         end1   = end2 
         }
    }
    result = append(result, [] int{start1, end1})
    return result
    
}