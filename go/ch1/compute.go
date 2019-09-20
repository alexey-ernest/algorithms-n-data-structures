package ch1

import "math"

func compute(s string) int {
	operators := new(StackLinkedList)
	values := new(StackLinkedList)

	for _, c := range s {
		switch c {
		case '(':
			continue
		case ' ':
			continue
		case '+':
			operators.Push(byte(c))
		case '-':
			operators.Push(byte(c))
		case '*':
			operators.Push(byte(c))
		case '/':
			operators.Push(byte(c))
		case '^':
			operators.Push(byte(c))
		case ')':
			val2 := int(values.Pop())
			op := operators.Pop()
			switch op {
			case '+':
				val := int(values.Pop())
				values.Push(byte(val + val2))
			case '-':
				val := int(values.Pop())
				values.Push(byte(val - val2))
			case '*':
				val := int(values.Pop())
				values.Push(byte(val * val2))
			case '/':
				val := int(values.Pop())
				values.Push(byte(val / val2))
			case '^':
				val := int(values.Pop())
				res := math.Pow(float64(val), float64(val2))
				values.Push(byte(res))
			}
		default:
			values.Push(byte(c) - '0')
		}
	}

	res := int(values.Pop())
	return res
}