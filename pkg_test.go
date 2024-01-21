// Package tests.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"fmt"
	"testing"
)

func TestReplaceCaseInsensitive(t *testing.T) {

	var tests = []struct {
		originalStr  string
		expectedStr  string
		originalWord string
		replacedWord string
	}{
		{
			"This is just a random string",
			"This is just a random string",
			"just",
			"just",
		},
		{
			"This is just a random string",
			"This is testing a random string",
			"just",
			"testing",
		},
		{
			"This is just a random string",
			"This is just a random string",
			"JUST",
			"just",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("replace_%s_with_%s", tt.originalWord, tt.replacedWord)
		t.Run(testname, func(t *testing.T) {

			ans := replaceCaseInsensitive(tt.originalStr, tt.originalWord, tt.replacedWord)

			if ans != tt.expectedStr {
				t.Errorf("Expected String(%s) is not same as"+
					" result string (%s)", tt.expectedStr, ans)
			}

		})
	}
}
