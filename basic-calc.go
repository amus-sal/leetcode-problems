// Stack
type (
	Stack struct {
		top *node
	}
	node struct {
		value int
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil}
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() int {
	n := this.top
	this.top = n.prev
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value int) {
	n := &node{value, this.top}
	this.top = n
}

func calculate(s string) int {
	// remove spaces
	s = strings.Replace(s, " ", "", -1)

	sum := 0
	num := 0
	multiplier := 1
	stack := New()
	for i := len(s) - 1; i >= 0; i-- {
		// push existing sum to stack
		// set sum to 0
		if s[i] == ')' {
			stack.Push(sum)
			sum = 0
			continue
		}

		// replace sum inside the parathesis a single number "num"
		// pop the existing sum from the stack
		if s[i] == '(' {
			num += sum
			sum = stack.Pop()
			continue
		}

		// add num to sum
		// restart multiplier and num
		if s[i] == '+' {
			sum += num
			multiplier = 1
			num = 0
			continue
		}

		// subtract num to sum
		// restart multiplier and num
		if s[i] == '-' {
			sum += -1 * num
			multiplier = 1
			num = 0
			continue
		}

		// else continue forming number
		// convert the rune to an int
		// multiply digit by power of 10
		digit := int(s[i] - '0')
		num += digit * multiplier
		multiplier = multiplier * 10
	}

	sum += num
	return sum
}