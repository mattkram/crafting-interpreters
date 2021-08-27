package main

import (
	"fmt"
	"os"
)

func runFile(filename string) error {
	fmt.Println("Running file:", filename)
	return nil
}

func runPrompt() error {
	fmt.Println("Running prompt")
	return nil
}

func runLoxMain() int {
	// Run the main Lox interpreter, either by passing in a filename or as a REPL
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		return 64
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return 64
		}
	} else {
		err := runPrompt()
		if err != nil {
			fmt.Println(err)
			return 64
		}
	}
	return 0
}

func main() {
	os.Exit(runLoxMain())
}
