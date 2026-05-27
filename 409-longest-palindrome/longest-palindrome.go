func longestPalindrome(s string) int {

    // Store frequency of each character
    freq := make(map[rune]int)

    // Stores the maximum palindrome length
    length := 0

    // Flag to check if any character has odd frequency
    hasOdd := false

    // Count frequency of each character
    for _, ch := range s {
        freq[ch]++
    }

    // Iterate through character frequencies
    for _, count := range freq {

        // If frequency is even,
        // we can use all occurrences
        if count%2 == 0 {
            length += count

        } else {

            // For odd frequency,
            // use count - 1 to make it even
            length += count - 1

            // Mark that we found an odd frequency
            hasOdd = true
        }
    }

    // If there is at least one odd frequency,
    // place one character in the center
    if hasOdd {
        length++
    }

    return length
}