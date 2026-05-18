func removeDuplicates(s string) string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}
