func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    var result [][]int;
    i,j := 0,0
    for i<len(firstList) && j< len(secondList){
        start1 := firstList[i][0]
        end1 := firstList[i][1]
        start2 := secondList[j][0]
        end2 := secondList[j][1]
        
        start := max(start1, start2)
        end := min(end1, end2)

        if start <= end {
            result = append(result, []int{start, end})
        }
        if (end1 <= end2){
            i+=1
        }else{
            j+=1
        }

    }
    return result
    
}