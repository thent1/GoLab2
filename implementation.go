package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// Stack implements a stack data structure.
type Stack struct {
	items []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil // Stack is empty
	}
	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Remove the top element
	return popped
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil // Stack is empty
	}
	return s.items[len(s.items)-1]
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the stack is empty, otherwise false
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// PostfixToInfix converts a postfix expression to an infix expression.
func PostfixToInfix(input string) (string, error) {

	if input == "" {
		return "", fmt.Errorf("Wrong expression, too many operands")
	}

	stack := Stack{}
	elements := strings.Split(input, " ")
	var buffer1 string
	var buffer2 string

	for i := 0; i < len(elements); i++ {
		if isValidOperand(elements[i]) {

			stack.Push(elements[i])

		} else if isOperator(elements[i]) {

			if stack.Size() > 1 {

				buffer2 = fmt.Sprintf("%v", stack.Pop())
				buffer1 = fmt.Sprintf("%v", stack.Pop())

				stack.Push(fmt.Sprintf("(%s %s %s)", buffer1, elements[i], buffer2))

			} else {
				return "", fmt.Errorf("Not enough operands in expression")
			}

		} else {
			return "", fmt.Errorf("Wrong symbols in expression")
		}
	}

	if stack.Size() == 1 {
		return fmt.Sprintf("%v", stack.Pop()), nil
	} else {
		return "", fmt.Errorf("Wrong expression, too many operands")
	}
}

// isOperator проверяет, является ли символ оператором
func isOperator(c string) bool {
	return len(c) == 1 && strings.ContainsAny(c, "+-*/^")
}

// isValidOperand checks if a string is a valid operand (number or letter)
func isValidOperand(c string) bool {
	_, err := strconv.Atoi(c)
	if err == nil {
		return true
	}

	if len(c) == 1 {
		r := rune(c[0])
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}

	return false
}
