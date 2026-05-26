func maxNumberOfBalloons(text string) int {
    //optimized version 
    // Target word
    word := "balloon"
    // Count characters available in text
    have := make(map[rune]int)

    // Count characters needed for word
    need := make(map[rune]int)

    for _, c := range text {
        have[c]++
    }

    for _, c := range word {
        need[c]++
    }

    // Large initial value
    ans := int(^uint(0) >> 1)

    // Find minimum possible formations
    for ch, freq := range need {

        possible := have[ch] / freq

        if possible < ans {
            ans = possible
        }
    }

    return ans
}