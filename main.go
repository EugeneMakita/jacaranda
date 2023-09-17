package main

import (
	"Lox/parser"
	"Lox/scanner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Jacaranta lang")
	for {
		reader := bufio.NewReader(os.Stdin)
		lines, _ := reader.ReadString('\n')
		lines = lines[:len(lines)-1]
		err := RunLine(lines)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func RunLine(lines string) error {
	if len(lines) <= 0 {
		return fmt.Errorf("line can't be empty")
	}

	scanner := scanner.CreateScanner(lines)
	tokens, err := scanner.CreateTokens()
	if err != nil {
		panic(err.Error())
	}

	for key, val := range tokens {
		fmt.Println(key, val)
	}

	parser := parser.CreateParser(tokens)
	p := parser.Equality()
	fmt.Println(p.String())

	return nil
}

func RunFile() error {
	return fmt.Errorf("placeholder")
}
