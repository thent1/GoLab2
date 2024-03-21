package lab2

import (
	"bufio"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	for scanner.Scan() {
		expression := scanner.Text()
		result, err := PostfixToInfix(strings.TrimSpace(expression))
		if err != nil {
			return err
		}
		if _, err := io.WriteString(ch.Output, result+"\n"); err != nil {
			return err
		}
	}
	return nil
}
