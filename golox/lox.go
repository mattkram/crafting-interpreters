package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func runString(source string) error {
	fmt.Println("Running string:", source)
	run(source)
	return nil
}

func runFile(filename string) error {
	fmt.Println("Running file:", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	run(string(b))
	return nil
}

func runPrompt() error {
	const inputPrefix = ">>> "
	fmt.Println("Running prompt")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(inputPrefix)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" || line == "quit" {
			return nil
		}
		if strings.TrimSpace(line) != "" {
			run(line)
		}
		fmt.Print(inputPrefix)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error encountered:", err)
		return err
	}
	return nil
}

func run(source string) {
	fmt.Println(source)
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

func main() {
	os.Exit(runLoxMain())
}
