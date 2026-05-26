func maxNumberOfBalloons(text string) int {

    // Target word
    word := "balloon"

    // Store character frequencies from text
    have := make(map[rune]int)

    // Store required frequencies for "balloon"
    need := make(map[rune]int)

    count := 0
    boolean := true

    // Count characters available in text
    for _, v := range text {
        have[v]++
    }

    // Count characters needed for "balloon"
    for _, v := range word {
        need[v]++
    }

    // Keep forming "balloon" until impossible
    for boolean {

        // Check every required character
        for k, v := range need {

            // Not enough characters available
            if v > have[k] || have[k] == 0 {
                boolean = false
                break
            }

            // Use required characters
            have[k] -= v
        }

        // Successfully formed one "balloon"
        if boolean {
            count++
        }
    }

    return count
}