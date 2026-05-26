func canConstruct(ransomNote string, magazine string) bool {

    // Create a hashmap to store frequency of characters
    m := make(map[rune]int)

    // Count each character in magazine
    for _, v := range magazine {
        m[v]++
    }

    // Try to use characters for ransomNote
    for _, v := range ransomNote {

        // Decrease count because character is used
        m[v]--

        // If count becomes negative,
        // magazine does not have enough characters
        if m[v] < 0 {
            return false
        }
    }

    // All characters were available
    return true
}