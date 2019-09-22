package ch1

func IsBalancedParentheses(s string) bool {
	var stack StackLinkedList
	
	for _, c := range s {
		chr := byte(c)

		switch chr {
		case '(', '[', '{':
			stack.Push(chr)
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}

			last := stack.Pop()
			
			switch chr {
			case ')':
				if last != '(' {
					return false
				}
			case ']':
				if last != '[' {
					return false
				}
			case '}':
				if last != '{' {
					return false
				}
			}
		}
	}

	return stack.IsEmpty()
}