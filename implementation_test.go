package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Invalid Expression",
			input:    "42 -",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Invalid Expression",
			input:    "!",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Invalid Expression",
			input:    "~",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "!",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "2 20 * 2 / 3 4 + 3 2 ^ * + 6 - 15 +",
			expected: "(((((2 * 20) / 2) + ((3 + 4) * (3 ^ 2))) - 6) + 15)",
			wantErr:  false,
		},
		{
			name:     "Valid Expression",
			input:    "4 2 - 3 * 5 +",
			expected: "(((4 - 2) * 3) + 5)",
			wantErr:  false,
		},
		{
			name:     "Valid Expression",
			input:    "3 15 + 4 * 6 4 - 1 7 + / - 3 * 16 +",
			expected: "(((((3 + 15) * 4) - ((6 - 4) / (1 + 7))) * 3) + 16)",
			wantErr:  false,
		},
		{
			name:     "Valid Expression",
			input:    "5 3 / 4 + 6 4 * 2 9 * + + 4 - 10 /",
			expected: "(((((5 / 3) + 4) + ((6 * 4) + (2 * 9))) - 4) / 10)",
			wantErr:  false,
		},
		{
			name:     "Expression with empty string",
			input:    "",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PostfixToInfix(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

// ExamplePostfixToInfix is a function that illustrating using of function PostfixToInfix
func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("4 2 - 3 * 5 +")
	fmt.Println(res)
	// Output: (((4 - 2) * 3) + 5)
}
