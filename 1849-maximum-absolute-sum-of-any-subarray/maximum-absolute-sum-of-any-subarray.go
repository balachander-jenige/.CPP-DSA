func maxAbsoluteSum(nums []int) int {
    maxEnding, minEnding := 0, 0
    res := 0

    for _, v := range nums {
        maxEnding = max(v, maxEnding+v)
        minEnding = min(v, minEnding+v)

        res = max(res, abs(maxEnding), abs(minEnding))
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}