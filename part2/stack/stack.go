package stack

type Stack struct {
	stack []string
}

func (r Stack) IsEmpty() bool {
	return len(r.stack) == 0
}

func (r *Stack) Push(value ...string) {
	r.stack = append(r.stack, value...)
}

func (r *Stack) Top() string {
	if r.IsEmpty() {
		return ""
	}

	return r.stack[len(r.stack)-1]
}

func (r *Stack) Pop() (value string) {
	if r.IsEmpty() {
		return ""
	}

	r.stack, value = r.stack[:len(r.stack)-1], r.stack[len(r.stack)-1]
	return
}
