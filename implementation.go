package lab2

import "fmt"

// TODO: document this function.
// PostfixToInfix returns postfix expression converted into infix expression
func PostfixToInfix(expression string) (string, error) {

	// Стек для збереження проміжних результатів
	stack := []string{}

	// Пройдемося по кожному символу виразу
	for _, char := range expression {
		// Якщо символ - це операція, видаляємо два останніх операнди зі стеку, формуємо інфіксний вираз
		if isOperator(string(char)) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: there are not enough operands for the operation")
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Перевірка на пріоритет операцій
			infixExpression := fmt.Sprintf("%s %s %s", operand1, string(char), operand2)
			if needsParentheses(string(char)) {
				infixExpression = fmt.Sprintf("(%s)", infixExpression)
			}
			stack = append(stack, infixExpression)
		} else if string(char) != " " { // Пропускаємо пробіл
			// Якщо символ - це операнд, просто додаємо його до стеку
			stack = append(stack, string(char))
		}
	}

	// Перевірка на кінцевий результат та наявність лишніх операндів
	if len(stack) != 1 {
		return "", fmt.Errorf("incorrect expression: extra operands")
	}

	return stack[0], nil
}

// needsParentheses determines if parentheses are needed around an operator.
// It returns true if the operator is either "+" or "-".
// Otherwise, it returns false.
func needsParentheses(operator string) bool {
	return operator == "+" || operator == "-"
}

// isOperator return true if input symbol is operator (+, -, *, /, ^) and return false otherwise
func isOperator(char string) bool {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
	_, ok := operators[char]
	return ok
}
