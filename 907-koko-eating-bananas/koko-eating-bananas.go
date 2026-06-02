func fun(piles []int, guess int)int {
    h := 0
    for _,v := range piles{
        h += v/guess
        if v% guess != 0{
            h++
        }
    }
    return h
}
func minEatingSpeed(piles []int, h int) int {
    low := 1
    high := slices.Max(piles)
    res := -1
    for low <= high{
        guess := (low+high)/2
        minhours := fun(piles,guess)
        if minhours > h{
            low = guess +1
        }else{
            res = guess
            high = guess-1
        }
    }
    return res
     
}