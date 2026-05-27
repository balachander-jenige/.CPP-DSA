func numberOfSpecialChars(word string) int {
    
    st := make(map[rune]bool)

    for _, ch := range word {
        st[ch] = true
    }

    count := 0

    for ch := 'a'; ch <= 'z'; ch++ {

        upper := ch - 'a' + 'A'

        if st[ch] && st[upper] {
            count++
        }
    }

    return count
}