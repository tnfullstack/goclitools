package main

import (
	"errors"
	"fmt"
)

// Errors declareation
var (
	ErrValidation = errors.New("validation failed")
)

// stepErr
type stepErr struct {
	step  string
	msg   string
	cause error
}

// Error
func (s *stepErr) Error() string {
	return fmt.Sprintf("step: %q: %s: cause: %v", s.step, s.msg, s.cause)
}

// Is
func (s *stepErr) Is(tg error) bool {
	t, ok := tg.(*stepErr)

	if !ok {
		return false
	}

	return t.step == s.step
}

// Unwrap
func (s *stepErr) Unwrap() error {
	return s.cause
}
