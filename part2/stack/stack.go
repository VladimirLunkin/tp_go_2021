package stack

import "errors"

type Stack struct {
	stack []string
}

const errorStackEmpty = "stack is empty"

func (r Stack) IsEmpty() bool {
	return len(r.stack) == 0
}

func (r *Stack) Push(value ...string) {
	r.stack = append(r.stack, value...)
}

func (r *Stack) Top() (string, error) {
	if r.IsEmpty() {
		return "", errors.New(errorStackEmpty)
	}

	return r.stack[len(r.stack)-1], nil
}

func (r *Stack) Pop() (value string, err error) {
	if r.IsEmpty() {
		return "", errors.New(errorStackEmpty)
	}

	r.stack, value = r.stack[:len(r.stack)-1], r.stack[len(r.stack)-1]
	return
}
