package ch1

func FixParentheses(input string) string {
	var stack StackLinkedListString
	for _, r := range input {
		b := string(r)
		
		switch b {
		case ")":
			var cmp string
			for range [3]int{} {
				cmp = stack.Pop() + cmp
			}

			// adding leading parentheses
			cmp = "(" + cmp + ")"
			stack.Push(cmp)

		default:
			stack.Push(b)
		}
	}

	return stack.Pop()
}