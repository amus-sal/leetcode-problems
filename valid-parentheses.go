


func isValid(s string) bool {
	var stack Stack

	for i := 0; i < len(s); i++ {

		if isClosedChar(string(s[i])) {
			cls, _ := stack.Pop()
			if !isMatched(string(s[i]), cls) {
				return false
			}
		} else {
			stack.Push(string(s[i]))
		}

	}

	return stack.IsEmpty()
}

func isMatched(st string, open string) bool {
	if st == ")" && open == "(" {
		return true
	}
	if st == "}" && open == "{" {
		return true
	}
	if st == "]" && open == "[" {
		return true
	}
	return false
}

func isClosedChar(st string) bool {
	if st == ")" || st == "}" || st == "]" {
		return true
	}

	return false
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

