package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((4 - 2) * 3) + 5)", res)
	}

	res, err = PostfixToInfix("abc++")
	if assert.Nil(t, err) {
		assert.Equal(t, "(a + (b + c))", res)
	}

	res, err = PostfixToInfix("ab*c+")
	if assert.Nil(t, err) {
		assert.Equal(t, "((a * b) + c)", res)
	}

	res, err = PostfixToInfix("abc/-ad/e-*")
	if assert.Nil(t, err) {
		assert.Equal(t, "((a - (b / c)) * ((a / d) - e))", res)
	}

	res, err = PostfixToInfix("23-4+567*+*")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((2 - 3) + 4) * (5 + (6 * 7)))", res)
	}

	res, err = PostfixToInfix("KL+MN*-OP^W*U/V/T*+Q+")
	if assert.Nil(t, err) {
		assert.Equal(t, "((((K + L) - (M * N)) + (((((O ^ P) * W) / U) / V) * T)) + Q)", res)
	}

	res, err = PostfixToInfix("ABCDEFGHOJ+-+-+-+//")
	if assert.Nil(t, err) {
		assert.Equal(t, "(A / (B / (C + (D - (E + (F - (G + (H - (O + J)))))))))", res)
	}

	res, err = PostfixToInfix("")
	if assert.NotNil(t, err) {
		assert.Equal(t, "", res)
	}

	res, err = PostfixToInfix("AB+-+-+-+-+-//^**")
	if assert.NotNil(t, err) {
		assert.Equal(t, "", res)
	}

	res, err = PostfixToInfix("ABOBA+ABOBA+ABOBA-ABOBA")
	if assert.NotNil(t, err) {
		assert.Equal(t, "", res)
	}
}

// ExamplePostfixToInfix is a function that illustrating using of function PostfixToInfix
func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("4 2 - 3 * 5 +")
	fmt.Println(res)
	// Output: (((4 - 2) * 3) + 5)
}
