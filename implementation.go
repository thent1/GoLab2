package lab2

import (
	"fmt"
	"strings"
)

// PostfixToInfix returns postfix expression converted into infix expression
func PostfixToInfix(expression string) (string, error) {
	// Стек для збереження проміжних результатів
	stack := []string{}

	// Функция isOperator проверяет, является ли символ оператором
	isOperator := func(char string) bool {
		operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
		_, ok := operators[char]
		return ok
	}

	// Разделяем выражение на операнды и операторы по пробелам
	tokens := strings.Split(expression, " ")

	// Проходим по каждому токену в выражении
	for _, token := range tokens {
		// Если токен - это оператор, извлекаем два последних операнда из стека
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("неправильний вираз: не вистачає операндів для операції")
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Формируем инфиксное выражение и добавляем его в стек
			infixExpression := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, infixExpression)
		} else { // Если токен - это операнд, добавляем его в стек
			stack = append(stack, token)
		}
	}

	// Проверяем наличие лишних операндов
	if len(stack) != 1 {
		return "", fmt.Errorf("неправильний вираз: лишні операнди")
	}

	return stack[0], nil
}
