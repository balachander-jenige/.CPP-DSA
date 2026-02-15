import "math"
func maximumSum(arr []int) int {

    nodelete := arr[0]
    onedelete := math.MinInt
    res := arr[0]

    for i:= 1; i<len(arr); i++{
        prevNoDelete := nodelete
        prevOneDelete := onedelete

        nodelete = max(prevNoDelete+arr[i],arr[i])

        if prevOneDelete == math.MinInt {
            onedelete = arr[i]
        }else{
            prevOneDelete = prevOneDelete +arr[i]
        }

        onedelete =max(prevOneDelete,prevNoDelete)
        res=max(res,nodelete,onedelete)
    }
    return res
}