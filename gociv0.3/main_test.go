package main

import (
	"bytes"
	"errors"
	"testing"
)

// TestRun
func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		proj   string
		out    string
		expErr error
	}{
		{
			name:   "Success",
			proj:   "./testdata/tool/",
			out:    "Go Build: SUCCESS\n GO Test: SUCCESS\n",
			expErr: nil,
		},
		{
			name:   "fail",
			proj:   "./testdata/toolErr",
			out:    "",
			expErr: &stepErr{step: "go build"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(tc.proj, &out)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error: %q. Got 'nil' instead.", tc.expErr)
					return
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("expected error: %q. Got %q.", tc.expErr, err)
				}
				return
			}
		})
	}
}
