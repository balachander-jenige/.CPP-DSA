func isValid(s string) bool {
	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// opening brackets
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			// stack empty -> invalid
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			// mismatch
			if top != pairs[ch] {
				return false
			}

			// pop
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}