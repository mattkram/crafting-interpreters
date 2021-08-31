package main

import (
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	// Store old OS arguments and delete on cleanup
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Create a temporary file such that run by filename will work
	os.Create("filename.lox")
	defer os.Remove("filename.lox")

	cases := []struct {
		Name         string
		Args         []string
		ExpectedExit int
	}{
		{"run by filename", []string{"filename.lox"}, 0},
		{"run by filename missing", []string{"missingfile.lox"}, 1},
		{"run with string arg", []string{"-c", "print \"Hello\""}, 0},
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

func TestError(t *testing.T) {
	// After emitting an error message the global HasError flag is set.
	if HadError {
		t.Errorf("Expected global error flag to start out as false")
	}
	Error(1, "message")
	if !HadError {
		t.Errorf("Expected global error flag to be true")
	}
}
