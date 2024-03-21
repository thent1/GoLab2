package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid Expression",
			input:    "42 1 -",
			expected: "(42 - 1)\n",
			wantErr:  false,
		},
		{
			name:     "Invalid Expression",
			input:    "42 -",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "5 5 +",
			expected: "(5 + 5)\n",
			wantErr:  false,
		},
		{
			name:     "Invalid Expression",
			input:    "-",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "3 4 * 5 +",
			expected: "((3 * 4) + 5)\n",
			wantErr:  false,
		},
		{
			name:     "Invalid Expression",
			input:    "1+",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "3 4 + 5 2 - *",
			expected: "((3 + 4) * (5 - 2))\n",
			wantErr:  false,
		},
		{
			name:     "Invalid Expression",
			input:    "a +",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid Expression",
			input:    "5 6 ^",
			expected: "(5 ^ 6)\n",
			wantErr:  false,
		},
		{
			name:     "Valid Expression",
			input:    "5 3 ^ 4 2 * - 6 1 + /",
			expected: "(((5 ^ 3) - (4 * 2)) / (6 + 1))\n",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			handler := ComputeHandler{
				Input:  input,
				Output: &output,
			}

			err := handler.Compute()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, output.String())
			}
		})
	}
}
