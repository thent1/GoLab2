package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func isOperator(c string) bool {
	return len(c) == 1 && strings.ContainsAny(c, "+-*/^")
}

func isValidOperand(c string) bool {
	_, err := strconv.Atoi(c)
	return err == nil
}

func PostfixToInfix(expression string) (string, error) {
	stack := []string{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if isOperator(token) || isValidOperand(token) {
			continue
		}
		return "", fmt.Errorf("invalid input: invalid character %s", token)
	}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid input: not enough operands for operator")
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			infixExpression := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, infixExpression)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("incorrect input: extra operands")
	}

	return stack[0], nil
}
