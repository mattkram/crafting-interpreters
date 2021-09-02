package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var HadError = false

func runString(source string) error {
	fmt.Println("Running string:", source)
	err := run(source)
	return err
}

func runFile(filename string) error {
	fmt.Println("Running file:", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = run(string(b))
	return err
}

func runPrompt() error {
	// Run as an interactive prompt (REPL).
	fmt.Println("Running prompt")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		if result := scanner.Scan(); !result {
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error encountered:", err)
			return err
		}

		line := strings.TrimSpace(scanner.Text())

		switch line {
		case "exit", "quit":
			return nil
		case "":
			continue
		}

		err := run(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func run(source string) error {
	s := NewScanner(source)
	tokens := s.scanTokens()

	for _, t := range tokens {
		fmt.Println(t.ToString())
	}

	return nil
}

func runLoxMain() int {
	// Run the main Lox interpreter, either by passing in a filename or as a REPL
	if len(os.Args) == 3 && os.Args[1] == "-c" {
		err := runString(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		return 64
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		err := runPrompt()
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}
	return 0
}

func Error(line int, msg string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error: %s\n", line, msg)
	HadError = true
}

func main() {
	os.Exit(runLoxMain())
}
