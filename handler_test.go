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
			expected: "42 1 -\n",
			wantErr:  false,
		},
		{
			name:     "Invalid Expression",
			input:    "42 -",
			expected: "",
			wantErr:  true,
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
