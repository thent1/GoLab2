package main

import (
	"bufio"
	"flag"
	"fmt"
	lab2 "github.com/thent1/GoLab2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression  = flag.String("e", "", "Expression to compute")
	fileExpression   = flag.String("f", "", "File with expression")
	outputExpression = flag.String("o", "", "File for result writing (unnecessary)")
)

func main() {

	flag.Parse()

	var reader io.Reader
	var writer io.Writer

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else if *fileExpression != "" {
		file, err := os.Open(*fileExpression)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		fmt.Fprintln(os.Stderr, "You must pass an expression with -e or a file with -f.")
		return
	}

	if *outputExpression != "" {
		file, err := os.Create(*outputExpression)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			return
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating expression: %v\n", err)
	}
}
