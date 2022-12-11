package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

type StatusError struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusError) Error() string {
	return se.Message
}

func (se StatusError) Unwrap() error {
	return se.Err
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("in main, wrappedErr:", wrappedErr)
		}
	}
}
