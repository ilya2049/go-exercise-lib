package main

import (
	"errors"
	"fmt"
)

type TimeoutError struct {
	timeout string
}

func (e *TimeoutError) Timeout() string {
	return e.timeout
}

func (e *TimeoutError) Error() string {
	return "timeout error: " + e.timeout
}

func main() {
	var err error = &TimeoutError{
		timeout: "1s",
	}

	err = fmt.Errorf("error: %w", err)

	var t interface {
		Timeout() string
	}

	if errors.As(err, &t) {
		fmt.Print(t.Timeout())
	}
}
