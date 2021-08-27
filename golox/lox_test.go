package main

import (
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		Name         string
		Args         []string
		ExpectedExit int
	}{
		{"run by filename", []string{"filename.lox"}, 0},
		{"run by prompt", []string{}, 0},
		{"too many args", []string{"f1.lox", "f2.lox"}, 64},
	}

	for _, tc := range cases {
		os.Args = append([]string{tc.Name}, tc.Args...)
		actualExit := runLoxMain()
		if tc.ExpectedExit != actualExit {
			t.Errorf("Wrong exit code for args: %v, expected %v, got %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}
	}
}
